package transaction

import (
	"context"
	"database/sql"
	"errors"
	"gotodo/app/repositories/connection"
	"gotodo/app/repositories/transactionctx"
	"log"
)

var ErrIsolationLevelMismatch = errors.New("transaction isolation level mismatch")

type Manager struct {
	connection *connection.DbConnection
}

func NewManager(connection *connection.DbConnection) *Manager {
	return &Manager{connection: connection}
}

func (tm *Manager) Begin(ctx context.Context, isolationLevel sql.IsolationLevel) (*transactionctx.TxContext, error) {
	tinfo, err := transactionctx.GetTransactionInfo(ctx)
	if err == nil {
		if tinfo.IsolationLevel != isolationLevel {
			return nil, ErrIsolationLevelMismatch
		}

		if !tinfo.IsDone() {
			return transactionctx.NewTxContext(ctx, tinfo), err
		}
	}

	tx, err := tm.connection.DB.BeginTx(ctx, &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tinfo = &transactionctx.Info{IsolationLevel: isolationLevel, Tx: tx}
	txctx := context.WithValue(ctx, transactionctx.TransactionInfoKey, tinfo)

	return transactionctx.NewTxContext(txctx, tinfo), nil
}

func (tm *Manager) BeginReadCommited(ctx context.Context) (*transactionctx.TxContext, error) {
	return tm.Begin(ctx, sql.LevelReadCommitted)
}
