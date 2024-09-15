package migration

import (
	"errors"
	"fmt"
	"transwallet/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB, log *logrus.Logger) {
	log.Info("Prepare for database table migration")
	err := db.AutoMigrate(&domain.Customer{}, &domain.Account{}, &domain.Transaction{})
	if err != nil {
		log.Error(err)
		fmt.Println(err)
	}
	log.Info("Database table migration has completed")
}

func CheckTableData[T any](db *gorm.DB, table T) error {

	// fetch table

	err := db.First(&table).Error
	if err != nil {
		// if no record return nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

	}
	return errors.New("table is not empty")
}

func AddDummyData(db *gorm.DB, log *logrus.Logger) {

	// Insert dummy customer
	log.Info("start inserting dummy data")
	err := CheckTableData[domain.Customer](db, domain.Customer{})
	if err == nil {
		customers := []domain.Customer{
			{ID: 1, Name: "Andi"},
			{ID: 2, Name: "Budi"},
			{ID: 3, Name: "Siti"},
		}
		for _, customer := range customers {
			db.Create(&customer)
		}
	} else {
		log.Error(err)
	}

	// Insert dummy account data
	err = CheckTableData[domain.Account](db, domain.Account{})
	if err == nil {
		accounts := []domain.Account{
			{ID: 100001, UserID: 1, Balance: 1000.0},
			{ID: 100002, UserID: 2, Balance: 1000.0},
			{ID: 100003, UserID: 3, Balance: 1000.0},
		}
		for _, account := range accounts {
			db.Create(&account)
		}
	} else {
		log.Error(err)
	}

	// Insert dummy transactions
	err = CheckTableData[domain.Transaction](db, domain.Transaction{})
	if err == nil {
		transactions := []domain.Transaction{
			{AccountID: 100001, Amount: 1000.0, Type: "c"},
			{AccountID: 100002, Amount: 1000.0, Type: "c"},
			{AccountID: 100003, Amount: 1000.0, Type: "c"},
		}
		for _, transaction := range transactions {
			db.Create(&transaction)
		}
	} else {
		log.Error(err)
	}
	log.Info("finish adding dummy data")

}
