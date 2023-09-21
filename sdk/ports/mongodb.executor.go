package ports

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	dbName string
}

func NewMongoExecutor(client *mongo.Client, dbName string) *Mongo {
	return &Mongo{
		client: client,
		dbName: dbName,
	}
}

func (ex *Mongo) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	return ex.client.StartSession(opts...)
}

func (ex *Mongo) Database(s ...types.DatabaseOpts) *mongo.Database {

	if len(s) > 0 {
		return ex.client.Database(s[0].Name, s[0].Opts...)
	}
	return ex.client.Database(ex.dbName)
}

func (ex *Mongo) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return ex.Database(types.DatabaseOpts{Name: ex.dbName}).Collection(name, opts...)
}

func (ex *Mongo) RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) (*mongo.Cursor, error) {
	return ex.Database(types.DatabaseOpts{Name: ex.dbName}).RunCommandCursor(ctx, runCommand, opts...)
}

func (ex *Mongo) RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) *mongo.SingleResult {
	return ex.Database(types.DatabaseOpts{Name: ex.dbName}).RunCommand(ctx, runCommand, opts...)
}

func (ex *Mongo) Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return ex.Database(types.DatabaseOpts{Name: ex.dbName}).Aggregate(ctx, pipeline, opts...)
}
