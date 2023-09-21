package conn

import (
	"database/sql"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var _ ports.IConnectionDB = (*MySql)(nil)

const (
	dbMysqlDriverDefault = "mysql"
)

type MySql struct {
	config   *types.DbConfig
	db       *sql.DB
	dbName   string
	connName string
}

// NewMySqlConnection constructor that returns a pointer to the IConnectionDB instance.
func NewMySqlConnection(config *types.DbConfig) *MySql {
	return &MySql{
		config: config,
	}
}

// BuildDNS function that constructs the DNS of the connection string to the selected database engine.
func (conn *MySql) BuildDNS() error {
	if conn.config == nil {
		return fmt.Errorf("Configuration file not found. ")
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

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		os.ExpandEnv(config.Username),
		os.ExpandEnv(config.Password),
		os.ExpandEnv(config.Host),
		os.ExpandEnv(strconv.Itoa(int(config.Port))),
		os.ExpandEnv(dbName),
		os.ExpandEnv(dbParams),
	)

	var dbDriver string
	if dbDriver = config.Driver; dbDriver == "" {
		dbDriver = dbMysqlDriverDefault
	}
	db, err := sql.Open(dbDriver, dataSourceName)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetConnMaxLifetime(config.ConnectionMaxLifetime.Duration)

	conn.db = db
	conn.dbName = dbName
	conn.connName = config.ConnectionName
	log.Printf("successfully connected to database %s", dbName)
	return nil
}

// Connect this function establishes a connection to the database credentials given to build the initial DNS.
func (conn *MySql) Connect() (any, error) {
	if conn.db == nil {
		return nil, fmt.Errorf("database is empty")
	}

	if err := conn.db.Ping(); err != nil {
		return nil, err
	}
	return conn.db, nil
}

// Migrate function that is in charge of performing the migrations in the database used in the DNS.
func (conn *MySql) Migrate() error {
	//TODO implementar esta funcion con go-migrate para las bases de datos mysql
	return nil
}

// Disconnect this function disconnects or terminates an existing connection to the database given in the DNS.
func (conn *MySql) Disconnect() error {
	if conn.db == nil {
		return fmt.Errorf("database is empty")
	}

	if err := conn.db.Close(); err != nil {
		return err
	}
	log.Printf("successfully disconnected to database %s", conn.dbName)
	return nil
}

// ClearConn this function is in charge of clearing the initial data given in the DNS, leaving the instance clean.
func (conn *MySql) ClearConn() error {
	if conn.db == nil {
		return fmt.Errorf("database is empty")
	}

	conn.db = nil
	conn.dbName = ""
	return nil
}

// GetDbName this function returns the name of the database in use.
func (conn *MySql) GetDbName() string {
	return conn.dbName
}
