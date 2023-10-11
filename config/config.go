package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type TomlConn struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

type TomlConfig struct {
	Conns []TomlConn
}

type Conn struct {
	Type string
	Dsn  string
}

type Config map[string]Conn

func New() *Config {
	tomlConfig := TomlConfig{}
	_, err := toml.DecodeFile("./config.toml", &tomlConfig)
	if err != nil {
		panic(err)
	}
	config := Config{}
	for _, v := range tomlConfig.Conns {
		switch v.Type {
		case "mysql":
			config[v.DB] = Conn{
				Type: v.Type,
				Dsn:  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", v.User, v.Password, v.Host, v.Port, v.DB),
			}
		case "postgres":
			config[v.DB] = Conn{
				Type: v.Type,
				Dsn:  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", v.Host, v.Port, v.User, v.Password, v.DB),
			}
		default:
			panic("invalid connection type")
		}
	}
	return &config
}

func (c *Config) ListDBs() []string {
	dbs := make([]string, 0, len(*c))
	for k := range *c {
		dbs = append(dbs, k)
	}
	return dbs
}

func (c *Config) GetConn(db string) (Conn, error) {
	conn, err := (*c)[db]
	if !err {
		return Conn{}, fmt.Errorf("invalid db: %s", db)
	}
	return conn, nil
}
