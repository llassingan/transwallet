package migration

import (
	"fmt"
	"transwallet/model/domain"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Account{}, &domain.Customer{}, &domain.Transaction{})
	if err != nil {
		fmt.Println(err)
	}
}

func AddDummyData(db *gorm.DB) {
	// Insert dummy account data
	accounts := []domain.Account{
		{ID: 100001, UserID: 1, Balance: 1000.0},
		{ID: 100002, UserID: 2, Balance: 1000.0},
		{ID: 100003, UserID: 3, Balance: 1000.0},
	}
	for _, account := range accounts {
		db.Create(&account)
	}

	// Insert dummy transactions
	transactions := []domain.Transaction{
		{ID: 1, AccountID: 100001, Amount: 1000.0,Type: "c"},
		{ID: 2, AccountID: 100002, Amount: 1000.0,Type: "c"},
		{ID: 3, AccountID: 100003, Amount: 1000.0,Type: "c"},
	}
	for _, transaction := range transactions {
		db.Create(&transaction)
	}

	// Insert dummy customer
	customers := []domain.Customer{
		{ID: 1, Name: "Andi"},
		{ID: 2, Name: "Budi"},
		{ID: 3, Name: "Siti"},
	}
	for _, customer := range customers {
		db.Create(&customer)
	}
}
