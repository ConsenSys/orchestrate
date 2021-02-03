package usecases

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store"
)

//go:generate mockgen -source=schedules.go -destination=mocks/schedules.go -package=mocks

/**
Schedule Use Cases
*/
type ScheduleUseCases interface {
	CreateSchedule() CreateScheduleUseCase
	GetSchedule() GetScheduleUseCase
	SearchSchedules() SearchSchedulesUseCase
}

type CreateScheduleUseCase interface {
	Execute(ctx context.Context, schedule *entities.Schedule) (*entities.Schedule, error)
	WithDBTransaction(dbtx store.Tx) CreateScheduleUseCase
}

type GetScheduleUseCase interface {
	Execute(ctx context.Context, scheduleUUID string, tenants []string) (*entities.Schedule, error)
}

type SearchSchedulesUseCase interface {
	Execute(ctx context.Context, tenants []string) ([]*entities.Schedule, error)
}