package db

import (
	"fmt"
	"time"

	"github.com/mrutman/simple-api/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/juju/loggo"
)

const (
	sqlDB     = "mysql"
	dbName    = "dbsimpleapi"
	tableName = "tablesimpleapi"
)

var (
	logger = loggo.GetLogger("db")
)

type SimpleRecord struct {
	ID        int       `db:"id"`
	Endpoint  string    `db:"endpoint"`
	Timestamp time.Time `db:"ts"`
}

func GetAllSimpleRecords() ([]SimpleRecord, error) {
	db, err := connectDB()
	if err != nil {
		logger.Errorf("Failed to get all SimpleRecord: '%v'", err)
		return nil, err
	}

	defer db.Close()

	all := []SimpleRecord{}
	err = db.Select(&all, fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		logger.Errorf("Failed to get all SimpleRecord: '%v'", err)
		return nil, err
	}

	return all, nil
}

func GetSimpleRecord(endpoint string) (*SimpleRecord, error) {
	db, err := connectDB()
	if err != nil {
		logger.Errorf("Failed to get SimpleRecord for endpoint '%s': '%v'", endpoint, err)
		return nil, err
	}

	defer db.Close()

	simpleRecord, err := getSimpleRecord(db, endpoint)
	if err != nil {
		return addSimpleRecord(db, endpoint)
	}

	return simpleRecord, nil
}

func connectDB() (*sqlx.DB, error) {
	cfg := config.GetSimpleAPIConfig()

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
		cfg.DBConnParam.User, cfg.DBConnParam.Password, cfg.DBConnParam.Host, cfg.DBConnParam.Port)
	db, err := sqlx.Open(sqlDB, connString)
	if err != nil {
		logger.Errorf("Failed to open DB with connection string '%s': '%v'", connString, err)
		return nil, err
	}

	if !cfg.DBParam.SkipCreateDB {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
		if err != nil {
			logger.Infof("Failed to create DB '%s': '%v'", dbName, err)
			db.Close()
			return nil, err
		}
	}

	_, err = db.Exec(fmt.Sprintf("USE %s", dbName))
	if err != nil {
		logger.Errorf("Failed to USE  DB '%s': '%v'", dbName, err)
		db.Close()
		return nil, err
	}

	if !cfg.DBParam.SkipCreateTables {
		schema := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY, `endpoint` varchar(255) NOT NULL, `ts` datetime)", tableName)
		db.Exec(schema)
		if err != nil {
			logger.Errorf("Failed to create table '%s' in DB '%s': '%v'", tableName, dbName, err)
			db.Close()
			return nil, err
		}
	}

	return db, nil
}

func getSimpleRecord(db *sqlx.DB, endpoint string) (*SimpleRecord, error) {
	simpleRecord := &SimpleRecord{}
	err := db.Get(simpleRecord, fmt.Sprintf("select * from %s where endpoint=\"%s\"", tableName, endpoint))
	if err != nil {
		logger.Errorf("Failed to get record for '%s' from table '%s': '%v'", endpoint, tableName, err)
		return nil, err
	}

	return simpleRecord, nil
}

func addSimpleRecord(db *sqlx.DB, endpoint string) (*SimpleRecord, error) {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO %s (endpoint, ts) VALUES('%s', NOW())", tableName, endpoint))
	if err != nil {
		logger.Errorf("Failed to add record for '%s' to table '%s': '%v'", endpoint, tableName, err)
		return nil, err
	}

	return getSimpleRecord(db, endpoint)
}
