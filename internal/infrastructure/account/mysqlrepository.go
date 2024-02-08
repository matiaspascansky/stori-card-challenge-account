package account

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"stori-card-challenge-account/domain/account"
)

type AccountDBRepository interface {
	SaveUserAccount(ctx context.Context, a *account.Account)
}

type accountDBRepository struct {
	db *sql.DB
}

func NewAccountDBRepository(db *sql.DB) *accountDBRepository {
	return &accountDBRepository{
		db: db,
	}
}

func (r *accountDBRepository) SaveUserAccount(ctx context.Context, a *account.Account) error {
	//aDto := FromAccountToDTO(a)
	//now := time.Now().UTC()

	_, err := r.db.ExecContext(ctx, "insert int")

	return err

}
