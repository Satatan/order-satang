package configs

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	App            appConfigs
	PostgresConfig PostgresConfig
}

type appConfigs struct {
	Port string `envconfig:"APP_PORT"`
	ENV  string `envconfig:"APP_ENV"`
}

var config Config

func Init() *Config {
	err := godotenv.Load()

	if err != nil {
		envFileNotFound := strings.Contains(err.Error(), "no such file or directory")
		if !envFileNotFound {
			logrus.WithFields(logrus.Fields{
				"path": "config/config.go",
				"func": "Init",
			}).Fatalf("read config err := %v", err)
		} else {
			logrus.Info("use environment from OS")
		}
	}
	if err = envconfig.Process("", &config); err != nil {
		logrus.WithFields(logrus.Fields{
			"path": "config/config.go",
			"func": "Init",
		}).Fatalf("parse configs error: %v", err)
	}
	return &config

}
