package domain

type Customer struct {
	ID   uint   `gorm:"primaryKey;column:id;autoIncrement"`
    Name string `gorm:"column:name"`
	AccountNumber Account `gorm:"foreignKey:UserID;references:ID"`
}

// set the table name 
func (u *Customer) TableName() string {
	return "customers"
}
