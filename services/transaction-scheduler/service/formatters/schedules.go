package formatters

import (
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/entities"
)

func FormatScheduleResponse(schedule *entities.Schedule) *types.ScheduleResponse {
	scheduleResponse := &types.ScheduleResponse{
		UUID:      schedule.UUID,
		TenantID:  schedule.TenantID,
		CreatedAt: schedule.CreatedAt,
		Jobs:      []*types.JobResponse{},
	}

	for idx := range schedule.Jobs {
		scheduleResponse.Jobs = append(scheduleResponse.Jobs, FormatJobResponse(schedule.Jobs[idx]))
	}

	return scheduleResponse
}
