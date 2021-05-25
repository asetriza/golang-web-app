package postgresql

import (
	"fmt"
	"strconv"
)

type Config struct {
	PostgreSQLUserName string
	PostgreSQLPassword string
	PostgreSQLHost     string
	PostgreSQLPort     int64
	PostgreSQLDBName   string
	SSLMode            string
}

// NewConfig creates new config of database, returns pointer to config or panics on error.
func NewConfig(user, password, host, port, dbName, sslMode string) (Config, error) {
	port16, err := strconv.ParseInt(port, 10, 16)
	if err != nil {
		return Config{}, err
	}

	return Config{
		PostgreSQLUserName: user,
		PostgreSQLPassword: password,
		PostgreSQLHost:     host,
		PostgreSQLPort:     port16,
		PostgreSQLDBName:   dbName,
		SSLMode:            sslMode,
	}, nil
}

// GetConnectionString creates connection string from config struct and returns string.
func (c Config) GetConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		c.PostgreSQLUserName,
		c.PostgreSQLPassword,
		c.PostgreSQLHost,
		c.PostgreSQLPort,
		c.PostgreSQLDBName,
		c.SSLMode)
}
