package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB{
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	// read config
	err := config.ReadInConfig()
	if err != nil{
		fmt.Println(err)
	}

	dsn := config.GetString("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println(err)
	}
	return db
}