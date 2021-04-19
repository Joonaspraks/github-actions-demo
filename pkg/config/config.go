package config

import (
	"fmt"
)

// Config hold the information in service configuration file
type Config struct {
	HTTP       *MainHTTPConfig   `config:"http"`
	PostgreSQL *PostgreSQLConfig `config:"postgresql"`
}

// PostgreSQLConfig holds configuration for PostgreSQL
type PostgreSQLConfig struct {
	Connection *PostgreSQLConnection `config:"connection"`
}

// URI builds postgres uri
func (p *PostgreSQLConfig) URI() string {
	return fmt.Sprintf(
		"dbname=%s host=%s password=%s user=%s sslmode=%s",
		p.Connection.DB,
		p.Connection.Host,
		p.Connection.Pass,
		p.Connection.User,
		p.Connection.SslMode,
	)
}

// PostgreSQLConnection holds configuration for PostgreSQL connection
type PostgreSQLConnection struct {
	Host    string `config:"host"`
	Port    int    `config:"port"`
	User    string `config:"user"`
	Pass    string `config:"pass"`
	DB      string `config:"db"`
	SslMode string `config:"ssl_mode"`
}

// MainHTTPConfig holds configuration for http setup of the application server
type MainHTTPConfig struct {
	Service *HTTPConfig `config:"service"`
}

// HTTPConfig holds configuration for a particular service
type HTTPConfig struct {
	Port uint   `config:"port"`
	Path string `config:"path"`
}
