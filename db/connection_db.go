package db

import (
	"fmt"
	"time"
	"transwallet/db/migration"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(log *logrus.Logger) *gorm.DB {
	// create viper instance
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".") // relative to project file

	// read config
	log.Info("Reading configuration")
	err := config.ReadInConfig()
	if err != nil {
		log.Error(err)
	}

	// get config 
	host := config.GetString("DB_HOST")
	port := config.GetInt("DB_PORT")
	user := config.GetString("DB_USER")
	dbname := config.GetString("DB_NAME")
	password := config.GetString("DB_PASSWORD")
	sslmode := config.GetString("DB_SSLMODE")
	
	// construct db connection
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s connect_timeout=5",
		host, port, user, dbname, password, sslmode)
	log.Info("Open connection")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	// set connection pool 
	sqlDB, err := db.DB()
	if err != nil {
		log.Error(err)
	}

	// configure connection pool
	maxOpenConns := config.GetInt("DB_MAX_OPEN_CONNS")
	maxIdleConns := config.GetInt("DB_MAX_IDLE_CONNS")
	connMaxLifetime,_ := time.ParseDuration(config.GetString("DB_CONN_MAX_LIFETIME"))

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	// start the migration 
	migration.Migration(db,log)
	migration.AddDummyData(db,log)
	return db
}
