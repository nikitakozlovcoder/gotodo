package transactionctx

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
)

const TransactionInfoKey = "transactionInfo"

var ErrTransactionInfoNotFound = errors.New("transaction info not found")

type TxContext struct {
	context.Context
	transactionInfo *Info
}

type Info struct {
	IsolationLevel sql.IsolationLevel
	Tx             *sql.Tx
}

func NewTxContext(ctx context.Context, transactionInfo *Info) *TxContext {
	return &TxContext{Context: ctx, transactionInfo: transactionInfo}
}

func GetTransactionInfo(ctx context.Context) (*Info, error) {
	info, ok := ctx.Value(TransactionInfoKey).(*Info)
	if !ok {
		return nil, ErrTransactionInfoNotFound
	}

	return info, nil
}

func (i *Info) IsDone() bool {
	isDone := reflect.ValueOf(i.Tx).MethodByName("isDone").Call([]reflect.Value{})
	return isDone[0].Bool()
}

func (t *TxContext) Rollback() {
	_ = t.transactionInfo.Tx.Rollback()
}

func (t *TxContext) Commit() error {
	if err := t.transactionInfo.Tx.Commit(); err != nil {
		return err
	}

	return nil
}
