package schedules

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/use-cases"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/log"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/parsers"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/models"
)

const getScheduleComponent = "use-cases.get-schedule"

// getScheduleUseCase is a use case to get a schedule
type getScheduleUseCase struct {
	db     store.DB
	logger *log.Logger
}

// NewGetScheduleUseCase creates a new GetScheduleUseCase
func NewGetScheduleUseCase(db store.DB) usecases.GetScheduleUseCase {
	return &getScheduleUseCase{
		db:     db,
		logger: log.NewLogger().SetComponent(getScheduleComponent),
	}
}

// Execute gets a schedule
func (uc *getScheduleUseCase) Execute(ctx context.Context, scheduleUUID string, tenants []string) (*entities.Schedule, error) {
	ctx = log.WithFields(ctx, log.Field("schedule", scheduleUUID))

	scheduleModel, err := fetchScheduleByUUID(ctx, uc.db, scheduleUUID, tenants)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(getScheduleComponent)
	}

	uc.logger.Debug("schedule found successfully")
	return parsers.NewScheduleEntityFromModels(scheduleModel), nil
}

func fetchScheduleByUUID(ctx context.Context, db store.DB, scheduleUUID string, tenants []string) (*models.Schedule, error) {
	schedule, err := db.Schedule().FindOneByUUID(ctx, scheduleUUID, tenants)
	if err != nil {
		return nil, err
	}

	for idx, job := range schedule.Jobs {
		schedule.Jobs[idx], err = db.Job().FindOneByUUID(ctx, job.UUID, tenants)
		if err != nil {
			return nil, err
		}
	}

	return schedule, nil
}