package usecases

import (
	"context"
	"fmt"
	"stori-card-challenge-account/domain/account"
	accountInfra "stori-card-challenge-account/internal/infrastructure/account"

	"github.com/pkg/errors"
)

type SaveAccountUsecase interface {
	Execute(ctx context.Context, account *account.Account) error
}

type saveAccountUsecase struct {
	accountRepository accountInfra.AccountDBRepository
}

func NewSaveAccountUsecase(accountRepository accountInfra.AccountDBRepository) *saveAccountUsecase {
	return &saveAccountUsecase{
		accountRepository: accountRepository,
	}
}

func (s *saveAccountUsecase) Execute(ctx context.Context, account *account.Account) error {
	err := validateModel(account)

	if err != nil {
		return errors.Wrap(err, "usecase: model is not defined correctly")
	}
	err = s.accountRepository.SaveUserAccount(ctx, account)

	if err != nil {
		return errors.Wrap(err, "usecase: cannot save user account")

	}
	return err
}

func validateModel(account *account.Account) error {
	if account.Id == "" {
		return fmt.Errorf("ID is required")
	}
	if account.DateCreated.IsZero() {
		return fmt.Errorf("DateCreated is required")
	}
	if account.Status == "" {
		return fmt.Errorf("Status is required")
	}

	return nil
}
