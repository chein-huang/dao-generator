
package dao

import (
	"context"
	"database/sql"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type transactionKey struct{}
type transactionInfo struct {
	tx   *gorm.DB
	err  error
	once sync.Once
}

func (d *dao) Transaction(ctx context.Context, f func(ctx context.Context) error, opts ...*sql.TxOptions) (err error) {
	txCtx := d.WithTransaction(ctx, opts...)
	defer func() {
		finishErr := d.FinishTransaction(txCtx)
		if finishErr != nil {
			err = errors.WithStack(finishErr)
		}
	}()

	err = f(txCtx)
	if err != nil {
		return ProcessError(txCtx, err)
	}
	return nil
}

func (d *dao) WithTransaction(ctx context.Context, opts ...*sql.TxOptions) context.Context {
	tx := d.db.WithContext(ctx).Begin(opts...)
	return context.WithValue(ctx, transactionKey{}, &transactionInfo{tx: tx})
}

func (d *dao) tryGetTx(ctx context.Context) *gorm.DB {
	v := ctx.Value(transactionKey{})
	if v != nil {
		if info, ok := v.(*transactionInfo); ok {
			return info.tx
		}
	}
	return d.db.WithContext(ctx)
}

func ProcessError(ctx context.Context, err error) error {
	v := ctx.Value(transactionKey{})
	if v != nil && err != nil {
		if info, ok := v.(*transactionInfo); ok {
			info.err = err
		}
	}
	return err
}

func (d *dao) FinishTransaction(ctx context.Context) error {
	v := ctx.Value(transactionKey{})
	if v != nil {
		if info, ok := v.(*transactionInfo); ok {
			var err error
			info.once.Do(func() {
				if info.err != nil {
					err = info.tx.Rollback().Error
				} else {
					err = info.tx.Commit().Error
				}
			})
			return err
		}
	}
	return nil
}