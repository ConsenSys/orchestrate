package parsers

import (
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/tx"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/entities"
)

func NewJobModelFromEntities(job *entities.Job, scheduleID *int) *models.Job {
	jobModel := &models.Job{
		UUID:       job.UUID,
		Type:       job.Type,
		Labels:     job.Labels,
		ScheduleID: scheduleID,
		Schedule: &models.Schedule{
			UUID: job.ScheduleUUID,
		},
		CreatedAt: job.CreatedAt,
	}

	if scheduleID != nil {
		jobModel.Schedule.ID = *scheduleID
	}

	if job.Transaction != nil {
		jobModel.Transaction = NewTransactionModelFromEntities(job.Transaction)
	}

	return jobModel
}

func NewJobEntityFromModels(jobModel *models.Job) *entities.Job {
	status := jobModel.GetStatus()
	if status == "" {
		status = entities.JobStatusCreated
	}

	job := &entities.Job{
		UUID:      jobModel.UUID,
		Status:    status,
		CreatedAt: jobModel.CreatedAt,
		Labels:    jobModel.Labels,
		Type:      jobModel.Type,
	}

	if jobModel.Schedule != nil {
		job.ScheduleUUID = jobModel.Schedule.UUID
	}

	if jobModel.Transaction != nil {
		job.Transaction = NewTransactionEntityFromModels(jobModel.Transaction)
	}

	return job
}

func UpdateJobModelFromEntities(jobModel *models.Job, job *entities.Job) {
	// for k, v := range job.Labels {
	// 	jobModel.Labels[k] = v
	// }
	// @TODO: Decide whether or not we should do a full replace (code above)
	if job.Labels != nil && len(job.Labels) > 0 {
		jobModel.Labels = job.Labels
	}

	UpdateTransactionModelFromEntities(jobModel.Transaction, job.Transaction)
}

func NewEnvelopeFromJobModel(job *models.Job, headers map[string]string) *tx.TxEnvelope {
	txEnvelope := &tx.TxEnvelope{
		Msg: &tx.TxEnvelope_TxRequest{TxRequest: &tx.TxRequest{
			Id:      job.UUID,
			JobUUID: job.UUID,
			JobType: tx.JobTypeMap[job.Type],
			Headers: headers,
			Params: &tx.Params{
				From:           job.Transaction.Sender,
				To:             job.Transaction.Recipient,
				Gas:            job.Transaction.GasLimit,
				GasPrice:       job.Transaction.GasPrice,
				Value:          job.Transaction.Value,
				Nonce:          job.Transaction.Nonce,
				Data:           job.Transaction.Data,
				Raw:            job.Transaction.Raw,
				PrivateFor:     job.Transaction.PrivateFor,
				PrivateFrom:    job.Transaction.PrivateFrom,
				PrivacyGroupId: job.Transaction.PrivacyGroupID,
			},
			ContextLabels: job.Labels,
		}},
		InternalLabels: make(map[string]string),
	}
	txEnvelope.SetChainUUID(job.Schedule.ChainUUID)

	return txEnvelope
}
