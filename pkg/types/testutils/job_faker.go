package testutils

import (
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"

	"github.com/gofrs/uuid"
)

func FakeJob() *entities.Job {
	return &entities.Job{
		UUID:         uuid.Must(uuid.NewV4()).String(),
		ScheduleUUID: uuid.Must(uuid.NewV4()).String(),
		ChainUUID:    uuid.Must(uuid.NewV4()).String(),
		Type:         utils.EthereumTransaction,
		InternalData: FakeInternalData(),
		Logs:         []*entities.Log{FakeLog()},
		CreatedAt:    time.Now(),
		Transaction:  FakeETHTransaction(),
	}
}

func FakeInternalData() *entities.InternalData {
	return &entities.InternalData{
		ChainID:       "888",
		Priority:      utils.PriorityMedium,
		RetryInterval: 5 * time.Second,
	}
}
