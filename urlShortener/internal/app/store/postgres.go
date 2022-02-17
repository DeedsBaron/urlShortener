package store

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

//func NewClient(ctx context.Context, username, password, host, port, database string) {
//	//dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
//		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
//		defer cancel()
//		//pool, err := pgxpool.Connect(ctx, dsn)
//		//if err != nil {
//		//	fmt.Println("failed to coonect to postrgersql")
//		//	return
//		//}
//	}
//
//}
