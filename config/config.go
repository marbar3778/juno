package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// Config defines all necessary juno configuration parameters.
type Config struct {
	RPCNode    string         `toml:"rpc_node"`
	ClientNode string         `toml:"client_node"`
	DB         DatabaseConfig `toml:"database"`
}

// DatabaseConfig defines all database connection configuration parameters.
type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     uint64 `toml:"port"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

func getEnv(key, defaultVal string) string {
	if value, exits := os.LookupEnv(key); exits {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func dbENV() Config {
	rpcNode := getEnv("RPC_NODE", "")
	clientNode := getEnv("CLIENT_NODE", "")
	host := getEnv("DBHOST", "")
	port := getEnvAsInt("DBPORT", 1)
	name := getEnv("DBNAME", "")
	user := getEnv("DBUSER", "")
	password := getEnv("DBPASSWORD", "")
	fmt.Println(rpcNode)

	dbConfig := DatabaseConfig{
		Host:     host,
		Port:     uint64(port),
		Name:     name,
		User:     user,
		Password: password,
	}

	config := Config{
		RPCNode:    rpcNode,
		ClientNode: clientNode,
		DB:         dbConfig,
	}
	fmt.Println(config)
	return config
}

// ParseConfig attempts to read and parse a Juno config from the given file path.
// An error reading or parsing the config results in a panic.
func ParseConfig(configPath string) Config {
	if configPath == "" {
		return dbENV()
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read config"))
	}

	var cfg Config
	if _, err := toml.Decode(string(configData), &cfg); err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode config"))
	}

	return cfg
}
