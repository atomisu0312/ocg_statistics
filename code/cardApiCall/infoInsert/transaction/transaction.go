package transaction

import (
	"context"
	"database/sql"

	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
)

type Tx struct {
	db *sql.DB
}

func NewTx(db *sql.DB) *Tx {
	return &Tx{db: db}
}

// fn func()の引数はrepositoriesという表記にしておけば、どのサービスからも使える
// そのためにはサービスにおいてrepositoriesという変数を用意する必要がある
func (tx *Tx) ExecTx(ctx context.Context, fn func(*sqlc_gen.Queries) error) error {
	// read-committedでトランザクションを開始する
	txx, err := tx.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	q := sqlc_gen.New(tx.db).WithTx(txx) //
	defer func() {
		//ここで外から渡した関数を実行する。
		if err == nil {
			err = txx.Commit()
		}
		if err != nil {
			err = txx.Rollback()
		}
	}()

	err = fn(q)
	if err != nil {
		return err
	}

	return nil
}

func (tx *Tx) ExecNonTx(ctx context.Context, fn func(*sqlc_gen.Queries) error) error {
	// トランザクションを張らずに実行
	q := sqlc_gen.New(tx.db) //
	err := fn(q)             //ここで外から渡した関数を実行する。
	if err != nil {
		return err
	}
	return nil
}
