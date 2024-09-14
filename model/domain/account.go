package domain

type Account struct {
	ID      uint    `gorm:"primaryKey;autoIncrement;column:id;default:100001"`
	UserID  uint    `gorm:"column:user_id"`
	Balance float64 `gorm:"column:balance;type:double precision;not null"`
	// has many
	Transaction []Transaction `gorm:"foreignKey:AccountID;references:ID"`
	// belongs to
	Customer *Customer `gorm:"foreignKey:UserID;references:ID"`
}

// set the table name
func (u *Account) TableName() string {
	return "accounts"
}
