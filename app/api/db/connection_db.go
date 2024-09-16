package db

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"transwallet/api/db/migration"
	"transwallet/api/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(log *logrus.Logger) *gorm.DB {
	// create viper instance
	// config := viper.New()
	// config.SetConfigFile("config.env")
	// config.AddConfigPath(".") // relative to project file

	// // read config
	// log.Info("Reading DB configuration")
	// err := config.ReadInConfig()
	// if err != nil {
	// 	log.Error(err)
	// }

	// // get config 
	// host := config.GetString("DB_HOST")
	// port := config.GetInt("DB_PORT")
	// user := config.GetString("POSTGRES_USER")
	// dbname := config.GetString("POSTGRES_DB")
	// password := config.GetString("POSTGRES_PASSWORD")
	// sslmode := config.GetString("DB_SSLMODE")

	// get config  from docker
	host := os.Getenv("DB_HOST")
	port,_ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	
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
	
	maxOpenConns,_ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	maxIdleConns,_ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	connMaxLifetime,_ := time.ParseDuration(os.Getenv("DB_CONN_MAX_LIFETIME"))

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	// start the migration
	err = migration.CheckTableData[domain.Customer](db,domain.Customer{})
	if err == nil{
		migration.Migration(db,log)
		migration.AddDummyData(db,log)
	}
	
	return db
}
