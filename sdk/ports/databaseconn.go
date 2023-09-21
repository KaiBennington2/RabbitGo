package ports

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IConnectionDB interface {
	BuildDNS() error
	Connect() (any, error)
	Migrate() error
	Disconnect() error
	ClearConn() error
	GetDbName() string
}

type IExecutorSqlDB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func GetSqlExec(connMap map[string]IConnectionDB, nameConn string) IExecutorSqlDB {
	db, ok := connMap[nameConn]
	if ok {
		d, _ := db.Connect()
		return d.(*sql.DB)
	}
	return nil
}

func GetConnections(
	fn func(motorName string, dbConfig types.DbConfig) IConnectionDB,
	dbSettings ...types.DbConfig,
) (map[string]IConnectionDB, error) {
	if dbSettings == nil {
		return nil, errors.New("No database configurations were found to initiate connections. ")
	}

	resp := make(map[string]IConnectionDB)
	var err error

	for _, config := range dbSettings {
		tempConn := fn(config.Driver, config)
		if err = tempConn.BuildDNS(); err != nil {
			return nil, err
		}
		if _, err = tempConn.Connect(); err != nil {
			return nil, err
		}

		resp[fmt.Sprintf("%s_conn", config.ConnectionName)] = tempConn
	}
	return resp, err
}

// ------------------------------------------------------------------------------------------------------ //

type IExecutorNoSqlDB interface {
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	Database(s ...types.DatabaseOpts) *mongo.Database
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
	RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) (*mongo.Cursor, error)
	RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) *mongo.SingleResult
	Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (*mongo.Cursor, error)
}

func GetNoSqlExec(connMap map[string]IConnectionDB, nameConn string) IExecutorNoSqlDB {
	db, ok := connMap[nameConn]
	if ok {
		d, _ := db.Connect()
		return NewMongoExecutor(d.(*mongo.Client), db.GetDbName())
	}
	return nil
}
