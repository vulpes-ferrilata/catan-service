package infrastructure

import (
	"os"

	"github.com/pkg/errors"
)

var (
	ErrEnvironmentVariableNotSet error = errors.New("environment variable is not set")
)

func NewConfig() (*Config, error) {
	config := new(Config)

	databaseConfig, err := newDatabaseConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	config.Database = databaseConfig

	return config, nil
}

type Config struct {
	Database *DatabaseConfig
}

func newDatabaseConfig() (*DatabaseConfig, error) {
	databaseConfig := new(DatabaseConfig)

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_HOST")
	}
	databaseConfig.Host = dbHost

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_PORT")
	}
	databaseConfig.Port = dbPort

	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_USERNAME")
	}
	databaseConfig.Username = dbUsername

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_PASSWORD")
	}
	databaseConfig.Password = dbPassword

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_NAME")
	}
	databaseConfig.Name = dbName

	return databaseConfig, nil
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}
