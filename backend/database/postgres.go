package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"order_satang/configs"
)

func NewPostgresConn(conf configs.PostgresConfig) CustomGorm {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		conf.Host, conf.Port, conf.User, conf.Database, conf.Pass)
	logrus.WithFields(logrus.Fields{}).Info("[CONFIG] repositories connection: ", strings.ReplaceAll(connectionString, conf.Pass, "********"))

	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Fatal("POSTGRES Failed to create connection Error: " + err.Error())
	}
	if conf.Debug {
		conn = conn.Debug()
	}

	sqlDB, err := conn.DB()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Fatal("POSTGRES Failed to DB definition Error: " + err.Error())
	}

	sqlDB.SetMaxIdleConns(conf.Maxidle)
	sqlDB.SetMaxOpenConns(conf.Maxopen)
	sqlDB.SetConnMaxLifetime(conf.Maxlifetime * time.Hour)

	return Wrap(conn)
}
