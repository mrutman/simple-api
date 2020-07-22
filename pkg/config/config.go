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

	DBParam *DBParameters `json:"dbParam" yaml:"dbParam"`
}

// DBConnectionParameters contains DB connection parameters
type DBConnectionParameters struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

// DBConnectionParameters contains DB connection parameters
type DBParameters struct {
	SkipCreateDB     bool `json:"skipCreateDB" yaml:"skipCreateDB"`
	SkipCreateTables bool `json:"skipCreateTables" yaml:"skipCreateTables"`
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
		DBParam: &DBParameters{
			SkipCreateDB:     os.Getenv("SIMPLEAPI_SKIP_CREATE_DB") == "true",
			SkipCreateTables: os.Getenv("SIMPLEAPI_SKIP_CREATE_TABLES") == "true",
		},
	}
}
