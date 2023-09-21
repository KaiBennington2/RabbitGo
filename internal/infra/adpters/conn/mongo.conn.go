package conn

import (
	"context"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var _ ports.IConnectionDB = (*Mongo)(nil)

const (
	dbMongoDriverDefault = "mongodb"
)

var ctx = context.Background()

type Mongo struct {
	config   *types.DbConfig
	db       *mongo.Client
	dbName   string
	connName string
}

func NewMongoConnection(config *types.DbConfig) *Mongo {
	return &Mongo{
		config: config,
	}
}

// BuildDNS function that constructs the DNS of the connection string to the selected database engine.
func (conn *Mongo) BuildDNS() error {
	if conn.config == nil {
		return fmt.Errorf("config is empty")
	}

	config := conn.config
	var dbName string
	if config.DbName != "" {
		dbName = os.ExpandEnv(config.DbName)
	} else {
		if dbName = config.DefaultSchema; dbName == "" {
			return fmt.Errorf("not found database name: %s", dbName)
		}
	}

	var dbParams string
	for key, val := range config.Params {
		dbParams += fmt.Sprintf("%s=%s&", key, val)
	}

	if dbParams != "" {
		dbParams = dbParams[0 : len(dbParams)-1]
	}

	var dbDriver string
	if dbDriver = config.Driver; dbDriver == "" {
		dbDriver = dbMongoDriverDefault
	}

	dataSourceName := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?%s",
		os.ExpandEnv(dbDriver),
		os.ExpandEnv(config.Username),
		os.ExpandEnv(config.Password),
		os.ExpandEnv(config.Host),
		os.ExpandEnv(strconv.Itoa(int(config.Port))),
		os.ExpandEnv(dbName),
		os.ExpandEnv(dbParams),
	)
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    os.ExpandEnv(dbName),
		Username:      os.ExpandEnv(config.Username),
		Password:      os.ExpandEnv(config.Password),
	}
	clientOptions := options.Client().ApplyURI(dataSourceName).SetAuth(credential)

	clientOptions.SetMaxPoolSize(uint64(config.MaxOpenConnections))
	duration := time.Duration(config.MaxIdleConnections) * time.Minute
	clientOptions.SetMaxConnIdleTime(duration)
	clientOptions.SetConnectTimeout(config.ConnectionMaxLifetime.Duration)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error establishing MongoDB connection: %v", err)
	}

	//set conns to sql.Db and dbName
	conn.db = client //.Database(dbName)
	conn.dbName = dbName
	conn.connName = config.ConnectionName
	log.Println(fmt.Sprintf("successfully connected to database %s", dbName))
	return nil
}

// Connect this function establishes a connection to the database credentials given to build the initial DNS.
func (conn *Mongo) Connect() (any, error) {
	if conn.db == nil {
		return nil, fmt.Errorf("database is empty")
	}

	if err := conn.db.Ping(ctx, &readpref.ReadPref{}); err != nil {
		return nil, err
	}
	return conn.db, nil
}

// Migrate function that is in charge of performing the migrations in the database used in the DNS.
func (conn *Mongo) Migrate() error {
	// TODO implementar esta funcion con flyway migrations para las bases de datos mongo
	return nil
}

// Disconnect this function disconnects or terminates an existing connection to the database given in the DNS.
func (conn *Mongo) Disconnect() error {
	if conn.db == nil {
		return fmt.Errorf("database is empty")
	}

	if err := conn.db.Disconnect(ctx); err != nil {
		return err
	}
	log.Printf("successfully disconnected to database %s", conn.dbName)
	return nil
}

// ClearConn this function is in charge of clearing the initial data given in the DNS, leaving the instance clean.
func (conn *Mongo) ClearConn() error {
	if conn.db == nil {
		return fmt.Errorf("database is empty")
	}

	conn.db = nil
	conn.dbName = ""
	return nil
}

// GetDbName this function returns the name of the database in use.
func (conn *Mongo) GetDbName() string {
	return conn.dbName
}
