package db

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(log *logrus.Logger) *gorm.DB {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	// read config
	log.Info("Reading configuration")
	err := config.ReadInConfig()
	if err != nil {
		log.Error(err)
	}

	dsn := config.GetString("DSN")
	log.Info("Open connection")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	return db
}
