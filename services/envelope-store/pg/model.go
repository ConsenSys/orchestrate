package pg

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/multitenancy"

	encoding "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/encoding/proto"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/envelope"
	evlpstore "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/envelope-store"
)

// EnvelopeModel represent elements in `envelopes` table
type EnvelopeModel struct {
	tableName struct{} `pg:"envelopes"` //nolint:unused,structcheck // reason

	// UUID technical identifier
	ID int32

	// Envelope Identifier
	EnvelopeID string

	// Tenant Identifier
	TenantID string

	// Ethereum info about transaction
	ChainID string
	TxHash  string

	// Envelope
	Envelope []byte

	// Status
	Status   string
	StoredAt time.Time
	SentAt   time.Time
	MinedAt  time.Time
	ErrorAt  time.Time
}

// StatusInfo returns a proto formatted StatusInfo object from a model
func (model *EnvelopeModel) StatusInfo() *evlpstore.StatusInfo {
	return &evlpstore.StatusInfo{
		Status:   model.StatusFormatted(),
		StoredAt: utils.TimeToPTimestamp(model.StoredAt),
		SentAt:   utils.TimeToPTimestamp(model.SentAt),
		MinedAt:  utils.TimeToPTimestamp(model.MinedAt),
		ErrorAt:  utils.TimeToPTimestamp(model.ErrorAt),
	}
}

// StatusFormatted returns a proto Status enum
func (model *EnvelopeModel) StatusFormatted() evlpstore.Status {
	status, ok := evlpstore.Status_value[strings.ToUpper(model.Status)]
	if !ok {
		panic(fmt.Sprintf("invalid status %q", model.Status))
	}
	return evlpstore.Status(status)
}

// UnmarshalEnvelope returns a proto envelope from model
func (model *EnvelopeModel) UnmarshalEnvelope(e *envelope.Envelope) error {
	return encoding.Unmarshal(model.Envelope, e)
}

func (model *EnvelopeModel) ToStatusResponse() (*evlpstore.StatusResponse, error) {
	return &evlpstore.StatusResponse{
		StatusInfo: model.StatusInfo(),
	}, nil
}

func (model *EnvelopeModel) ToStoreResponse() (*evlpstore.StoreResponse, error) {
	resp := &evlpstore.StoreResponse{
		StatusInfo: model.StatusInfo(),
		Envelope:   &envelope.Envelope{},
	}

	// Unmarshal envelope
	err := model.UnmarshalEnvelope(resp.GetEnvelope())
	if err != nil {
		return &evlpstore.StoreResponse{}, err
	}

	return resp, nil
}

// FromEnvelope creates a model from an envelope
func FromEnvelope(ctx context.Context, e *envelope.Envelope) (*EnvelopeModel, error) {
	// Marshal envelope
	b, err := encoding.Marshal(e)
	if err != nil {
		return nil, err
	}

	tenantID := multitenancy.TenantIDFromContext(ctx)

	return &EnvelopeModel{
		Envelope:   b,
		EnvelopeID: e.GetMetadata().GetId(),
		TenantID:   tenantID,
		ChainID:    e.GetChain().GetBigChainID().String(),
		TxHash:     e.GetTx().GetHash().Hex(),
	}, nil
}
