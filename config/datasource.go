package config

import (
	"fmt"

	"github.com/imhshekhar47/go-api-core/utils"
)

type DatasourceConfig struct {
	Host     string `yaml:"host",mapstructure:"host"`
	Port     string `yaml:"port",mapstructure:"port"`
	Database string `yaml:"database",mapstructure:"database"`
	Username string `yaml:"username",mapstructure:"username"`
	Password string `yaml:"password",mapstructure:"password"`
}

func GetDatasourceConfig() DatasourceConfig {
	return DatasourceConfig{
		Host:     utils.GetEnvOrElse("DB_HOST", "localhost"),
		Port:     utils.GetEnvOrElse("DB_PORT", "0"),
		Database: utils.GetEnvOrElse("DB_NAME", ""),
		Username: utils.GetEnvOrElse("DB_USERNAME", ""),
		Password: utils.GetEnvOrElse("DB_PASSWORD", ""),
	}
}

func (s *DatasourceConfig) IsValid() error {
	if utils.IsEmpty(s.Database) {
		return fmt.Errorf("reqired field database is missing")
	}

	if utils.IsEmpty(s.Host) {
		return fmt.Errorf("required field host is missing")
	}

	if utils.IsEmpty(s.Port) {
		return fmt.Errorf("required field port is missing")
	}

	if utils.IsEmpty(s.Username) {
		return fmt.Errorf("required fiel username is empty")
	}

	if utils.IsEmpty(s.Password) {
		return fmt.Errorf("required field password is empty")
	}

	return nil
}
