package accounts

import (
	"context"

	parsers2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/parsers"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/use-cases"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/log"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store"
)

const updateAccountComponent = "use-cases.update-account"

type updateAccountUseCase struct {
	db     store.DB
	logger *log.Logger
}

func NewUpdateAccountUseCase(db store.DB) usecases.UpdateAccountUseCase {
	return &updateAccountUseCase{
		db:     db,
		logger: log.NewLogger().SetComponent(updateAccountComponent),
	}
}

func (uc *updateAccountUseCase) Execute(ctx context.Context, account *entities.Account, tenants []string) (*entities.Account, error) {
	ctx = log.WithFields(ctx, log.Field("address", account.Address))
	logger := uc.logger.WithContext(ctx)

	model, err := uc.db.Account().FindOneByAddress(ctx, account.Address, tenants)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(updateAccountComponent)
	}

	if account.Attributes != nil {
		model.Attributes = account.Attributes
	}
	if account.Alias != "" {
		model.Alias = account.Alias
	}

	err = uc.db.Account().Update(ctx, model)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(updateAccountComponent)
	}

	resp := parsers2.NewAccountEntityFromModels(model)

	logger.Info("account updated successfully")
	return resp, nil
}