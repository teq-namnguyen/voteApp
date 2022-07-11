package utils

import (
	"context"
	"gorm.io/gorm"
	"strings"
)

type KeyContext string

var (
	keyEnable    KeyContext = "postgresql_tx_enable"
	keyTx        KeyContext = "postgresql_tx"
	keyGetClient KeyContext = "db_get_client"
)

func TxBegin(ctx context.Context, getClient func(ctx context.Context) *gorm.DB) context.Context {
	db := getClient(ctx)
	tx := db.Begin()
	ctx = SetTx(ctx, tx)

	ctx = context.WithValue(ctx, keyEnable, true)
	return ctx
}

func TxEnd(ctx context.Context, txFunc func(context.Context) error) (context.Context, error) {
	var err error

	tx := GetTx(ctx)

	func(ctx context.Context) {
		defer func() {
			p := recover()

			switch {
			case p != nil:
				tx.Rollback()
				panic(p)
			case err != nil:
				tx.Rollback()
			default:
				err = tx.Commit().Error
				if err != nil {
					tx.Logger.Error(ctx, "fail commit transaction", err)
				}
			}

		}()
		err = txFunc(ctx)
	}(ctx)

	ctx = context.WithValue(ctx, keyEnable, false)

	return ctx, err
}

func IsEnableTx(ctx context.Context) bool {
	txEnable, ok := ctx.Value(keyEnable).(bool)

	return ok && txEnable
}

func GetTx(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(keyTx).(*gorm.DB)
	if !ok {
		return nil
	}

	return tx
}

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, keyTx, tx)
}

func GetGetClientFunc(ctx context.Context) func(ctx context.Context) *gorm.DB {
	getClient, ok := ctx.Value(keyGetClient).(func(ctx context.Context) *gorm.DB)
	if !ok {
		return nil
	}

	return getClient
}

func IsDeadlockError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "deadlock")
}
