package steps

import (
	"fmt"
	gohttp "net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/cucumber/godog"
	gherkin "github.com/cucumber/messages-go/v10"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/broker/sarama"
	encoding "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/encoding/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/tx"
	chainregistry "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/client"
	contractregistry "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/client"
	registry "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/proto"
	envelopestore "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/envelope-store/client"
	txscheduler "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/client"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/tests/service/chanregistry"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/tests/service/cucumber/parser"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/tests/service/cucumber/tracker"
)

const GenericNamespace = "_"

var TOPICS = [...]string{
	"tx.crafter",
	"tx.signer",
	"tx.sender",
	"tx.decoded",
	"tx.recover",
	"account.generator",
	"account.generated",
}

// AuthSetup is container for authentication context data
type AuthSetup struct {
	authMethod string
	authData   string
}

// ScenarioContext is container for scenario context data
type ScenarioContext struct {
	Pickle *gherkin.Pickle

	// Parser to parse cucumber/gherkin entries
	parser *parser.Parser

	// trackers track envelopes that are generated within the test session
	// as they are processed in the system
	trackers []*tracker.Tracker

	// defaultTracker allows to capture envelopes that are generated
	// within the system (to be captured those envelopes should have scenario.id set)
	defaultTracker *tracker.Tracker

	// chanReg to register envelopes channels on trackers
	chanReg *chanregistry.ChanRegistry

	httpClient   *gohttp.Client
	httpResponse *gohttp.Response

	aliases *parser.AliasRegistry

	// RegistryClient
	ContractRegistry registry.ContractRegistryClient

	// Transaction Schedule
	TransactionScheduler txscheduler.TransactionSchedulerClient

	// Producer to producer envelopes in topics
	producer sarama.SyncProducer

	logger *log.Entry

	authSetup AuthSetup
}

func setServiceURL(sc *ScenarioContext) {
	sc.aliases.Set(GenericNamespace, "chain-registry", viper.GetString(chainregistry.ChainRegistryURLViperKey))
	sc.aliases.Set(GenericNamespace, "chain-registry-metrics", viper.GetString(chainregistry.ChainRegistryMetricsURLViperKey))
	sc.aliases.Set(GenericNamespace, "contract-registry", viper.GetString(contractregistry.ContractRegistryURLViperKey))
	sc.aliases.Set(GenericNamespace, "contract-registry-metrics", viper.GetString(contractregistry.ContractRegistryMetricsURLViperKey))
	sc.aliases.Set(GenericNamespace, "contract-registry-http", viper.GetString(contractregistry.ContractRegistryHTTPURLViperKey))
	sc.aliases.Set(GenericNamespace, "envelope-store", viper.GetString(envelopestore.EnvelopeStoreURLViperKey))
	sc.aliases.Set(GenericNamespace, "envelope-store-metrics", viper.GetString(envelopestore.EnvelopeStoreMetricsURLViperKey))
	sc.aliases.Set(GenericNamespace, "envelope-store-http", viper.GetString(envelopestore.EnvelopeStoreHTTPURLViperKey))
	sc.aliases.Set(GenericNamespace, "tx-scheduler-metrics", viper.GetString(txscheduler.TxSchedulerMetricsURLViperKey))
	sc.aliases.Set(GenericNamespace, "tx-scheduler-http", viper.GetString(txscheduler.TxSchedulerURLViperKey))
}

func NewScenarioContext(
	chanReg *chanregistry.ChanRegistry,
	httpClient *gohttp.Client,
	contractRegistry registry.ContractRegistryClient,
	txScheduler txscheduler.TransactionSchedulerClient,
	producer sarama.SyncProducer,
	p *parser.Parser,
) *ScenarioContext {
	sc := &ScenarioContext{
		chanReg:              chanReg,
		httpClient:           httpClient,
		aliases:              p.GetAliasRegistry(),
		ContractRegistry:     contractRegistry,
		TransactionScheduler: txScheduler,
		producer:             producer,
		parser:               p,
		logger:               log.NewEntry(log.StandardLogger()),
		authSetup:            AuthSetup{},
	}

	setServiceURL(sc)

	return sc
}

// initScenarioContext initialize a scenario context - create a random scenario id - initialize a logger enrich with the scenario name - initialize envelope chan
func (sc *ScenarioContext) init(s *gherkin.Pickle) {
	// Hook the Pickle to the scenario context
	sc.Pickle = s

	// Prepare default tracker
	sc.defaultTracker = sc.newTracker(nil)

	// Enrich logger
	sc.logger = sc.logger.WithFields(log.Fields{
		"scenario.name": sc.Pickle.Name,
		"scenario.id":   sc.Pickle.Id,
	})

	sc.logger.Debug("new scenario initialized")
	sc.aliases.Set(sc.Pickle.Id, "SCENARIO_ID", sc.Pickle.Id)
}

func (sc *ScenarioContext) newTracker(e *tx.Envelope) *tracker.Tracker {
	// Create tracker and attach envelope
	t := tracker.NewTracker()
	t.Current = e

	// Initialize output channels on tracker and register channels on channel registry
	for _, topic := range TOPICS {
		// Create channel
		// TODO: make chan size configurable
		ch := make(chan *tx.Envelope, 30)

		// Add channel as a tracker output
		t.AddOutput(topic, ch)

		// Register channel on channel registry
		if e != nil {
			log.WithFields(log.Fields{
				"id":          e.GetID(),
				"scenario.id": sc.Pickle.Id,
				"topic":       topic,
			}).Debugf("registered new envelope")
			sc.chanReg.Register(LongKeyOf(topic, sc.Pickle.Id, e.GetID()), ch)
		} else {
			sc.chanReg.Register(ShortKeyOf(topic, sc.Pickle.Id), ch)
		}
	}

	return t
}

func (sc *ScenarioContext) setMetadata(e *tx.Envelope) {
	_ = e.SetID(uuid.Must(uuid.NewV4()).String())

	// Prepare envelope metadata
	_ = e.SetContextLabelsValue("debug", "true").
		SetContextLabelsValue("scenario.id", sc.Pickle.Id).
		SetContextLabelsValue("scenario.name", sc.Pickle.Name)
}

func (sc *ScenarioContext) newTrackers(envelopes []*tx.Envelope) []*tracker.Tracker {
	// Create a tracker for every envelope
	var trackers []*tracker.Tracker
	for _, e := range envelopes {
		// Create a tracker
		sc.setMetadata(e)
		trackers = append(trackers, sc.newTracker(e))
	}

	return trackers
}

func (sc *ScenarioContext) setTrackers(trackers []*tracker.Tracker) {
	sc.trackers = trackers
}

func (sc *ScenarioContext) sendEnvelope(topic string, e *tx.Envelope) error {
	// Prepare message to be sent
	msg := &sarama.ProducerMessage{
		Topic: viper.GetString(fmt.Sprintf("topic.%v", topic)),
		Key:   sarama.StringEncoder(e.KafkaPartitionKey()),
	}

	err := encoding.Marshal(e.TxEnvelopeAsRequest(), msg)
	if err != nil {
		return err
	}

	// Send message
	_, _, err = sc.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"id":            e.GetID(),
		"scenario.id":   sc.Pickle.Id,
		"scenario.name": sc.Pickle.Name,
	}).Debugf("scenario: envelope sent")

	return nil
}

func (sc *ScenarioContext) iSendEnvelopesToTopic(topic string, table *gherkin.PickleStepArgument_PickleTable) error {
	// Parse table
	envelopes, err := sc.parser.ParseEnvelopes(sc.Pickle.Id, table)
	if err != nil {
		return errors.DataError("could not parse tx request - got %v", err)
	}

	// Set trackers for each envelope
	sc.setTrackers(sc.newTrackers(envelopes))

	// Send envelopes
	for _, t := range sc.trackers {
		err := sc.sendEnvelope(topic, t.Current)
		if err != nil {
			return errors.InternalError("could not send tx request - got %v", err)
		}
	}

	return nil
}

func (sc *ScenarioContext) registerEnvelopeTracker(value string) error {
	envelopeID, ok := sc.aliases.Get(sc.Pickle.Id, value)
	if !ok {
		envelopeID, ok = sc.aliases.Get("global", value)
		if !ok {
			envelopeID = value
		}
	}

	evlp := tx.NewEnvelope()
	_ = evlp.SetID(envelopeID).
		SetContextLabelsValue("debug", "true").
		SetContextLabelsValue("scenario.id", sc.Pickle.Id)

	sc.setTrackers(append(sc.trackers, sc.newTracker(evlp)))

	return nil
}

func (sc *ScenarioContext) iHaveDeployedContract(alias string, table *gherkin.PickleStepArgument_PickleTable) error {
	// Parse table
	envelopes, err := sc.parser.ParseEnvelopes(sc.Pickle.Id, table)
	if err != nil {
		return err
	}

	// Set trackers
	trackers := sc.newTrackers(envelopes)

	if len(trackers) != 1 {
		return fmt.Errorf("%v: should deploy exactly 1 contract", sc.Pickle.Id)
	}

	// Send envelope
	err = sc.sendEnvelope("tx.crafter", trackers[0].Current)
	if err != nil {
		return err
	}

	// Catch envelope after it has been decoded
	err = trackers[0].Load("tx.decoded", 30*time.Second)
	if err != nil {
		return fmt.Errorf("%v: no receipt for contract %q deployment", sc.Pickle.Id, alias)
	}

	// Alias contract address
	if trackers[0].Current.GetReceipt().GetContractAddress() == "" {
		return fmt.Errorf("%v: contract %q could not be deployed", sc.Pickle.Id, alias)
	}
	sc.parser.Aliases.Set(sc.Pickle.Id, alias, trackers[0].Current.GetReceipt().GetContractAddress())

	return nil
}

func (sc *ScenarioContext) envelopeShouldBeInTopic(topic string) error {
	for i, t := range sc.trackers {
		err := t.Load(topic, viper.GetDuration(CucumberTimeoutViperKey))
		if err != nil {
			return fmt.Errorf("%v: envelope n°%v not in topic %q", sc.Pickle.Id, i, topic)
		}
	}
	return nil
}

func (sc *ScenarioContext) envelopesShouldContainTheFollowingErrors(table *gherkin.PickleStepArgument_PickleTable) error {
	rowsEnvelopes := table.Rows[1:]

	if len(rowsEnvelopes) != len(sc.trackers) {
		return fmt.Errorf("expected as much errors as envelopes tracked")
	}

	for r, t := range sc.trackers {
		row := rowsEnvelopes[r]
		if len(row.Cells) != len(t.Current.Errors) {
			return fmt.Errorf("(%d/%d) got %d errors but expected %d - got: %v", r+1, len(sc.trackers), len(t.Current.Errors), len(row.Cells), t.Current.Errors)
		}
		for i, err := range t.Current.Errors {
			if !strings.Contains(err.Message, row.Cells[i].Value) {
				return fmt.Errorf(
					"(%d/%d) got '%s' error but expected '%s'", r+1, len(sc.trackers), err.Message, row.Cells[i].Value)
			}
		}
	}
	return nil
}

func (sc *ScenarioContext) envelopesShouldHaveTheFollowingValues(table *gherkin.PickleStepArgument_PickleTable) error {
	header := table.Rows[0]
	rowsEnvelopes := table.Rows[1:]
	if len(rowsEnvelopes) != len(sc.trackers) {
		return fmt.Errorf("expected as much rows as envelopes tracked")
	}

	for r, row := range rowsEnvelopes {
		val := reflect.ValueOf(sc.trackers[r].Current).Elem()
		for c, col := range row.Cells {
			fieldName := header.Cells[c].Value
			field, err := getField(fieldName, val)
			if err != nil {
				return err
			}
			if col.Value == "~" {
				if isEqual("", field) {
					return fmt.Errorf("(%d/%d) did not expected %s to be empty", r+1, len(rowsEnvelopes), fmt.Sprintf("%v", fieldName))
				}
				continue
			}

			var aliasRE = regexp.MustCompile(`{{(.*)}}`)
			if aliasRE.MatchString(col.Value) {
				alias := aliasRE.FindStringSubmatch(col.Value)[1]
				val, _ := sc.aliases.Get(sc.Pickle.Id, alias)
				if !isEqual(val, field) {
					return fmt.Errorf("(%d/%d) %s expected %s but got %s", r+1, len(rowsEnvelopes), fieldName, val, fmt.Sprintf("%v", field))
				}

				continue
			}

			if !isEqual(col.Value, field) {
				return fmt.Errorf("(%d/%d) %s expected %s but got %s", r+1, len(rowsEnvelopes), fieldName, col.Value, fmt.Sprintf("%v", field))
			}
		}
	}
	return nil
}

func isEqual(s string, val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, _ := strconv.ParseInt(s, 10, 64)
		if val.Int() != n {
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, _ := strconv.ParseUint(s, 10, 64)
		if val.Uint() != n {
			return false
		}
	case reflect.Float32, reflect.Float64:
		n, _ := strconv.ParseFloat(s, 64)
		if val.Float() != n {
			return false
		}
	default:
		if val.String() != s && fmt.Sprintf("%v", val) != s {
			return false
		}
	}
	return true
}

func getField(fieldName string, val reflect.Value) (reflect.Value, error) {
	fieldName = strings.Title(fieldName)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if strings.Contains(fieldName, ".") {
		keyValue := strings.Split(fieldName, ".")

		field := val.FieldByName(keyValue[0])
		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		switch field.Kind() {
		case reflect.Slice, reflect.Array:
			i, _ := strconv.ParseInt(keyValue[1], 10, 64)
			if int(i) < field.Len() {
				return reflect.Value{}, fmt.Errorf("%s length is only %d could not reach %d", keyValue[0], field.Len(), i)
			}
			field = field.Index(int(i))
		case reflect.Map:
			field = field.MapIndex(reflect.ValueOf(strings.Title(keyValue[1])))
		default:
			field = field.FieldByName(keyValue[1])
		}

		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		if len(keyValue) > 2 {
			return getField(strings.Join(keyValue[2:], "."), field)
		}

		return field, nil
	}

	field := val.FieldByName(fieldName)

	return field, nil
}

func (sc *ScenarioContext) envelopesShouldHavePayloadSet() error {
	for i, t := range sc.trackers {
		if t.Current.GetData() == "" {
			return fmt.Errorf("%v: payload not set envelope n°%v", sc.Pickle.Id, i)
		}
	}
	return nil
}

func (sc *ScenarioContext) envelopesShouldHaveNonceSet() error {
	nonces := make(map[string]map[string]map[uint64]bool)
	for _, t := range sc.trackers {
		chain := t.Current.GetChainID().String()
		addr := t.Current.GetFromString()
		nonce, err := t.Current.GetNonceUint64()
		if err != nil {
			return err
		}

		if _, ok := nonces[chain]; !ok {
			nonces[chain] = make(map[string]map[uint64]bool)
		}
		if _, ok := nonces[chain][addr]; !ok {
			nonces[chain][addr] = make(map[uint64]bool)
		}
		if _, ok := nonces[chain][addr][nonce]; ok {
			return fmt.Errorf("%v: nonce %d attributed more than once", sc.Pickle.Id, t.Current.Nonce)
		}
		nonces[chain][addr][nonce] = true
	}

	return nil
}

func (sc *ScenarioContext) envelopesShouldHaveRawAndHashSet() error {
	for i, t := range sc.trackers {
		if t.Current.Raw == "" {
			return fmt.Errorf("%v: raw not set on envelope n°%v", sc.Pickle.Id, i)
		}

		if t.Current.TxHash == nil {
			return fmt.Errorf("%v: hash not set on envelope n°%v", sc.Pickle.Id, i)
		}
	}
	return nil
}

func (sc *ScenarioContext) envelopesShouldHaveFromSet() error {
	for i, t := range sc.trackers {
		if t.Current.From == nil {
			return fmt.Errorf("%v: from not set on envelope n°%v", sc.Pickle.Id, i)
		}
	}
	return nil
}

func (sc *ScenarioContext) envelopesShouldHaveLogDecoded() error {
	for i, t := range sc.trackers {
		for _, l := range t.Current.GetReceipt().GetLogs() {
			if len(l.GetTopics()) > 0 && len(l.GetDecodedData()) == 0 {
				return fmt.Errorf("%v: log have not been decoded on envelope n°%v", sc.Pickle.Id, i)
			}
		}
	}

	return nil
}

// FeatureContext is a initializer for cucumber scenario methods
func FeatureContext(s *godog.Suite) {
	sc := NewScenarioContext(
		chanregistry.GlobalChanRegistry(),
		http.NewClient(),
		contractregistry.GlobalClient(),
		txscheduler.GlobalClient(),
		broker.GlobalSyncProducer(),
		parser.GlobalParser(),
	)

	s.BeforeScenario(sc.init)

	s.BeforeStep(func(s *gherkin.Pickle_PickleStep) {
		log.WithFields(log.Fields{
			"step.text":     s.Text,
			"scenario.name": sc.Pickle.Name,
			"scenario.id":   sc.Pickle.Id,
		}).Debugf("scenario: step starts")
	})
	s.AfterStep(func(s *gherkin.Pickle_PickleStep, err error) {
		log.WithError(err).
			WithFields(log.Fields{
				"step.text":     s.Text,
				"scenario.name": sc.Pickle.Name,
				"scenario.id":   sc.Pickle.Id,
			}).Debugf("scenario: step completed")
	})

	initHTTP(s, sc)
	registerContractRegistrySteps(s, sc)

	s.Step(`^I have deployed contract "([^"]*)"$`, sc.iHaveDeployedContract)
	s.Step(`^I send envelopes to topic "([^"]*)"$`, sc.iSendEnvelopesToTopic)
	s.Step(`^Register new envelope tracker "([^"]*)"$`, sc.registerEnvelopeTracker)
	s.Step(`^Envelopes should be in topic "([^"]*)"$`, sc.envelopeShouldBeInTopic)
	s.Step(`^Envelopes should have payload set$`, sc.envelopesShouldHavePayloadSet)
	s.Step(`^Envelopes should have nonce set$`, sc.envelopesShouldHaveNonceSet)
	s.Step(`^Envelopes should have raw and hash set$`, sc.envelopesShouldHaveRawAndHashSet)
	s.Step(`^Envelopes should have from set$`, sc.envelopesShouldHaveFromSet)
	s.Step(`^Envelopes should have log decoded$`, sc.envelopesShouldHaveLogDecoded)
	s.Step(`^Envelopes should have the following fields:$`, sc.envelopesShouldHaveTheFollowingValues)
	s.Step(`^Envelopes should have the following errors:$`, sc.envelopesShouldContainTheFollowingErrors)
}
