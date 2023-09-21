package types

type FileExtType string

const (
	Yaml FileExtType = "yaml"
	Json FileExtType = "json"
	Env  FileExtType = "env"
)

type Setting struct {
	ServerConfig  SvrConfig       `yaml:"server_info" json:"server_info" xml:"server_info" env:"server_info"`
	DbConfig      []DbConfig      `yaml:"databases" json:"databases" xml:"databases" env:"databases"`
	EncryptConfig []EncryptConfig `yaml:"encryptors" json:"encryptors" xml:"encryptors" env:"encryptors"`
	BrokerConfig  []BrokerConfig  `yaml:"message_brokers" json:"message_brokers" xml:"message_brokers" env:"message_brokers"`
}

type DbConfig struct {
	ConnectionName        string            `yaml:"connection_name" json:"connection_name" xml:"connection_name" env:"connection_name"`
	Driver                string            `yaml:"driver" json:"driver" xml:"driver" env:"driver"`
	DbName                string            `yaml:"db_name" json:"db_name" xml:"db_name" env:"db_name"`
	Username              string            `yaml:"username" json:"username" xml:"username" env:"username"`
	Password              string            `yaml:"password" json:"password" xml:"password" env:"password"`
	Host                  string            `yaml:"host" json:"host" xml:"host" env:"host"`
	Port                  int64             `yaml:"port" json:"port" xml:"port" env:"port"`
	Params                map[string]string `yaml:"params" json:"params" xml:"params" env:"params"`
	MaxOpenConnections    int               `yaml:"max_open_connections" json:"max_open_connections" xml:"max_open_connections" env:"max_open_connections"`
	MaxIdleConnections    int               `yaml:"max_idle_connections" json:"max_idle_connections" xml:"max_idle_connections" env:"max_idle_connections"`
	ConnectionMaxLifetime Duration          `yaml:"connection_max_lifetime" json:"connection_max_lifetime" xml:"connection_max_lifetime" env:"connection_max_lifetime"`
	DefaultSchema         string            `yaml:"default_schema" json:"default_schema" xml:"default_schema" env:"default_schema"`
}

type SvrConfig struct {
	Port     int64   `yaml:"port" json:"port" xml:"port" env:"port"`
	PathBase *bool   `yaml:"path_base" json:"path_base" xml:"path_base" env:"path_base"`
	PathName *string `yaml:"path_name" json:"path_name" xml:"path_name" env:"path_name"`
}

// EncryptConfig Structure containing the credentials needed to configure the encryption tools.
type EncryptConfig struct {
	Name      string `yaml:"name" json:"name" xml:"name" env:"name"`
	Key       string `yaml:"key" json:"key" xml:"key" env:"key"`
	Algorithm string `yaml:"algorithm" json:"algorithm" xml:"algorithm" env:"algorithm"`
}

// BrokerConfig Structure containing the credentials needed to configure the messaging broker
type BrokerConfig struct {
	BrokerName     string            `yaml:"broker_name" json:"broker_name" xml:"broker_name" env:"broker_name"`
	ConnectionName string            `yaml:"connection_name" json:"connection_name" xml:"connection_name" env:"connection_name"`
	Driver         string            `yaml:"driver" json:"driver" xml:"driver" env:"driver"`
	Username       string            `yaml:"username" json:"username" xml:"username" env:"username"`
	Password       string            `yaml:"password" json:"password" xml:"password" env:"password"`
	Host           string            `yaml:"host" json:"host" xml:"host" env:"host"`
	Port           int64             `yaml:"port" json:"port" xml:"port" env:"port"`
	Params         map[string]string `yaml:"params" json:"params" xml:"params" env:"params"`
}
