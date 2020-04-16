// +build integration

package integrationtests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/abi"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/testutils"
	registry "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/proto"
)

type contractRegistryHTTPTestSuite struct {
	suite.Suite
	httpClient http.Client
	baseURL    string
	env        *IntegrationEnvironment
}

func (s *contractRegistryHTTPTestSuite) SetupSuite() {
	s.httpClient = http.Client{
		Timeout: 5 * time.Second,
	}
}

func (s *contractRegistryHTTPTestSuite) TestContractRegistry_Validation() {
	// TODO: Next test is returning 500
	// s.T().Run("should fail with 400 if payload is invalid", func(t *testing.T) {
	// 	contract := testutils.FakeContract()
	// 	contract.SetName("")

	// 	resp := s.registerContract(contract)
	// 	assert.Equal(t, 400, resp.StatusCode)
	// })

	s.T().Run("should not fail if contract registered twice", func(t *testing.T) {
		contract := testutils.FakeContract()

		resp := s.registerContract(contract)
		assert.Equal(t, 200, resp.StatusCode)

		resp = s.registerContract(contract)
		assert.Equal(t, 200, resp.StatusCode)
	})
}

func (s *contractRegistryHTTPTestSuite) TestContractRegistry_Register() {
	s.T().Run("should register a contract with tag", func(t *testing.T) {
		contract := testutils.FakeContract()
		contract.SetTag("tag")

		resp := s.registerContract(contract)
		assert.Equal(t, 200, resp.StatusCode)

		resp = s.getContract(contract.GetName(), contract.GetTag())
		assert.Equal(t, 200, resp.StatusCode)

		body := &registry.GetContractResponse{}
		getValue(resp, body)

		assert.Equal(t, body.GetContract().GetName(), contract.GetName())
		assert.Equal(t, body.GetContract().GetTag(), contract.GetTag())
	})

	s.T().Run("should register a contract with tag latest", func(t *testing.T) {
		contract := testutils.FakeContract()
		contract.SetTag("")

		resp := s.registerContract(contract)
		assert.Equal(t, 200, resp.StatusCode)

		resp = s.getContract(contract.GetName(), contract.GetTag())
		assert.Equal(t, 200, resp.StatusCode)

		body := &registry.GetContractResponse{}
		getValue(resp, body)

		assert.Equal(t, body.GetContract().GetName(), contract.GetName())
		assert.Equal(t, body.GetContract().GetTag(), "latest")
	})
}

func (s *contractRegistryHTTPTestSuite) TestContractRegistry_Get() {
	contract0 := testutils.FakeContract()
	_ = contract0.CompactABI()
	resp := s.registerContract(contract0)
	assert.Equal(s.T(), 200, resp.StatusCode)

	contract1 := testutils.FakeContract()
	resp = s.registerContract(contract1)
	assert.Equal(s.T(), 200, resp.StatusCode)

	s.T().Run("should get all contracts", func(t *testing.T) {
		resp = s.getCatalog()
		assert.Equal(t, 200, resp.StatusCode)
	})

	s.T().Run("should get all tags of a contract", func(t *testing.T) {
		resp = s.getTags(contract0.GetName())
		assert.Equal(t, 200, resp.StatusCode)

		body := &registry.GetTagsResponse{}
		getValue(resp, body)

		assert.Equal(t, body.GetTags(), []string{"v1.0.0"})
	})

	s.T().Run("should get a contract", func(t *testing.T) {
		resp = s.getContract(contract0.GetName(), contract0.GetTag())
		assert.Equal(t, resp.StatusCode, 200)

		body := &registry.GetContractResponse{}
		getValue(resp, body)

		assert.Equal(t, body.GetContract().GetName(), contract0.GetName())
		assert.Equal(t, body.GetContract().GetTag(), contract0.GetTag())
		assert.Equal(t, body.GetContract().GetBytecode(), contract0.GetBytecode())
		assert.Equal(t, body.GetContract().GetDeployedBytecode(), contract0.GetDeployedBytecode())
		assert.Equal(t, body.GetContract().GetAbi(), contract0.GetAbi())
	})

	s.T().Run("should get a contract abi", func(t *testing.T) {
		resp = s.getContractField(contract0.GetName(), contract0.GetTag(), "abi")
		assert.Equal(t, 200, resp.StatusCode)

		body := &registry.GetContractABIResponse{}
		getValue(resp, body)

		assert.Equal(t, body.GetAbi(), contract0.GetAbi())
	})

	s.T().Run("should get a contract bytecode", func(t *testing.T) {
		resp = s.getContractField(contract0.GetName(), contract0.GetTag(), "bytecode")
		assert.Equal(t, 200, resp.StatusCode)

		body := &registry.GetContractBytecodeResponse{}
		getValue(resp, body)

		assert.Equal(t, body.GetBytecode(), contract0.GetBytecode())
	})

	s.T().Run("should get a contract deployed bytecode", func(t *testing.T) {
		resp = s.getContractField(contract0.GetName(), contract0.GetTag(), "deployedBytecode")
		assert.Equal(t, 200, resp.StatusCode)

		body := &registry.GetContractDeployedBytecodeResponse{}
		getValue(resp, body)

		assert.Equal(t, body.GetDeployedBytecode(), contract0.GetDeployedBytecode())
	})
}

func (s *contractRegistryHTTPTestSuite) registerContract(contract *abi.Contract) *http.Response {
	request, err := json.Marshal(registry.RegisterContractRequest{
		Contract: contract,
	})
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(s.baseURL+"/contracts", "application/json", bytes.NewBuffer(request))
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *contractRegistryHTTPTestSuite) getCatalog() *http.Response {
	resp, err := http.Get(s.baseURL + "/contracts")
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *contractRegistryHTTPTestSuite) getContract(name, tag string) *http.Response {
	resp, err := http.Get(s.baseURL + "/contracts/" + name + "/" + tag)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *contractRegistryHTTPTestSuite) getContractField(name, tag, field string) *http.Response {
	resp, err := http.Get(s.baseURL + "/contracts/" + name + "/" + tag + "/" + field)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *contractRegistryHTTPTestSuite) getTags(name string) *http.Response {
	resp, err := http.Get(s.baseURL + "/contracts/" + name)
	if err != nil {
		panic(err)
	}

	return resp
}

func getValue(resp *http.Response, body interface{}) {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(err)
	}
}