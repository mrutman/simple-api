package config

import (
	"os"
)

const (
	defaultServerPort = "8888"
)

// SimpleAPIConfig contains configuration parameters for SimpleAPI
type SimpleAPIConfig struct {
	ServerPort string `json:"serverPort" yaml:"serverPort"`

	DBConnParam *DBConnectionParameters `json:"dbConnParam" yaml:"dbConnParam"`
}

// DBConnectionParameters contains DB connection parameters
type DBConnectionParameters struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

func GetSimpleAPIConfig() *SimpleAPIConfig {
	return GetSimpleAPIConfigFromEnv()
}

func GetSimpleAPIConfigFromEnv() *SimpleAPIConfig {
	serverPort := os.Getenv("SIMPLEAPI_SERVER_PORT")
	if len(serverPort) == 0 {
		serverPort = defaultServerPort
	}

	return &SimpleAPIConfig{
		ServerPort: serverPort,
		DBConnParam: &DBConnectionParameters{
			Host:     os.Getenv("SIMPLEAPI_DB_HOST"),
			Port:     os.Getenv("SIMPLEAPI_DB_PORT"),
			User:     os.Getenv("SIMPLEAPI_DB_USER"),
			Password: os.Getenv("SIMPLEAPI_DB_PASSWORD"),
		},
	}
}
