package test

import (
	"context"
	"errors"
	"stori-card-challenge-account/domain/account"
	"stori-card-challenge-account/domain/user"
	"stori-card-challenge-account/internal/infrastructure/account/mocks"
	usecases "stori-card-challenge-account/internal/usecases/account"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SaveAccountTestSuite struct {
	suite.Suite
	ctx         context.Context
	accountRepo *mocks.AccountDBRepository
	session     *session.Session
}

func (s *SaveAccountTestSuite) SetupTest() {

	s.ctx = context.TODO()
	s.accountRepo = new(mocks.AccountDBRepository)

}

func (s *SaveAccountTestSuite) TearDownTest() {
	s.accountRepo.AssertExpectations(s.T())
}

func TestSaveAccount(t *testing.T) {
	suite.Run(t, new(SaveAccountTestSuite))
}

func (s *SaveAccountTestSuite) Test_Save_Account() {
	s.T().Run("success", func(t *testing.T) {

		usr := user.User{
			ID:        1,
			FirstName: "Mati",
			LastName:  "Pas",
		}

		mockAccount := &account.Account{
			Id:           "testid",
			User:         usr,
			Status:       "testStatus",
			TotalBalance: 100.32,
			DateCreated:  time.Now().UTC(),
		}

		s.accountRepo.On("SaveUserAccount", s.ctx, mockAccount).Return(nil)

		saveAccountUsecase := usecases.NewSaveAccountUsecase(s.accountRepo)

		err := saveAccountUsecase.Execute(s.ctx, mockAccount)

		assert.NoError(t, err)
		s.accountRepo.AssertNumberOfCalls(t, "SaveUserAccount", 1)

	})

	s.TearDownTest()
	s.SetupTest()

	s.T().Run("error_missing_id", func(t *testing.T) {

		usr := user.User{
			ID:        1,
			FirstName: "Mati",
			LastName:  "Pas",
		}

		mockAccount := &account.Account{
			Id:           "testid",
			User:         usr,
			Status:       "testStatus",
			TotalBalance: 100.32,
		}

		saveAccountUsecase := usecases.NewSaveAccountUsecase(s.accountRepo)

		err := saveAccountUsecase.Execute(s.ctx, mockAccount)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "DateCreated is required")
		s.accountRepo.AssertNumberOfCalls(t, "SaveUserAccount", 0)

	})

	s.T().Run("error_missing_date", func(t *testing.T) {

		usr := user.User{
			ID:        1,
			FirstName: "Mati",
			LastName:  "Pas",
		}

		mockAccount := &account.Account{
			User:         usr,
			Status:       "testStatus",
			TotalBalance: 100.32,
			DateCreated:  time.Now().UTC(),
		}

		saveAccountUsecase := usecases.NewSaveAccountUsecase(s.accountRepo)

		err := saveAccountUsecase.Execute(s.ctx, mockAccount)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "ID is required")
		s.accountRepo.AssertNumberOfCalls(t, "SaveUserAccount", 0)

	})

	s.T().Run("error_missing_status", func(t *testing.T) {

		usr := user.User{
			ID:        1,
			FirstName: "Mati",
			LastName:  "Pas",
		}

		mockAccount := &account.Account{
			Id:           "testID",
			User:         usr,
			TotalBalance: 100.32,
			DateCreated:  time.Now().UTC(),
		}

		saveAccountUsecase := usecases.NewSaveAccountUsecase(s.accountRepo)

		err := saveAccountUsecase.Execute(s.ctx, mockAccount)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "Status is required")
		s.accountRepo.AssertNumberOfCalls(t, "SaveUserAccount", 0)

	})
	s.TearDownTest()
	s.SetupTest()
	s.T().Run("error_on_save", func(t *testing.T) {

		usr := user.User{
			ID:        1,
			FirstName: "Mati",
			LastName:  "Pas",
		}

		mockAccount := &account.Account{
			Id:           "testid",
			User:         usr,
			Status:       "testStatus",
			TotalBalance: 100.32,
			DateCreated:  time.Now().UTC(),
		}

		s.accountRepo.On("SaveUserAccount", s.ctx, mockAccount).Return(errors.New("foo"))

		saveAccountUsecase := usecases.NewSaveAccountUsecase(s.accountRepo)

		err := saveAccountUsecase.Execute(s.ctx, mockAccount)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "usecase: cannot save user account")
		s.accountRepo.AssertNumberOfCalls(t, "SaveUserAccount", 1)

	})

}
