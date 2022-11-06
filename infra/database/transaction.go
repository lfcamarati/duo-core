package database

import (
	"database/sql"
	"errors"
)

// https://betterprogramming.pub/how-to-write-atomic-repositories-in-go-f4fdef01f769

var (
	ErrTransactionInProgress = errors.New("transaction already in progress")
	ErrTransactionNotStarted = errors.New("transaction not started")
)

type Transactor interface {
	Begin() error
	Commit() error
	Rollback() error
}

type GenericTransactor struct {
	Db *sql.DB
	Tx *sql.Tx
}

func (r *GenericTransactor) Begin() error {
	if r.Tx != nil {
		return ErrTransactionInProgress
	}

	tx, err := r.Db.Begin()

	if err != nil {
		return err
	}

	r.Tx = tx
	return nil
}

func (r *GenericTransactor) Commit() error {
	if r.Tx == nil {
		return ErrTransactionNotStarted
	}

	err := r.Tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (r *GenericTransactor) Rollback() error {
	if r.Tx == nil {
		return ErrTransactionNotStarted
	}

	err := r.Tx.Rollback()

	if err != nil {
		return err
	}

	return nil
}
