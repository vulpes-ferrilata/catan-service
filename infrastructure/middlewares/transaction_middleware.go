package middlewares

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	infrastructure_context "github.com/VulpesFerrilata/catan-service/infrastructure/context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewTransactionMiddleware(db *gorm.DB) *TransactionMiddleware {
	return &TransactionMiddleware{
		db: db,
	}
}

type TransactionMiddleware struct {
	db *gorm.DB
}

func (t TransactionMiddleware) WrapHandler(next bus.HandlerFunc) bus.HandlerFunc {
	return func(ctx context.Context, input interface{}) (result interface{}, err error) {
		tx := t.db.WithContext(ctx).Begin()

		defer func() {
			if r := recover(); r != nil {
				if terr := tx.Rollback().Error; terr != nil {
					err = errors.WithStack(terr)
				}

				panic(r)
			}
		}()

		ctx = infrastructure_context.WithTransaction(ctx, tx)

		result, err = next(ctx, input)
		if err != nil {
			if terr := tx.Rollback().Error; terr != nil {
				err = errors.WithStack(terr)
			}
			return
		}

		if terr := tx.Commit().Error; terr != nil {
			err = errors.WithStack(terr)
		}

		return
	}
}
