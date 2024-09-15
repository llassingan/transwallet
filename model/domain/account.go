package domain

import "time"

type Account struct {
	ID        int      `gorm:"primaryKey;column:id"`
	UserID    int      `gorm:"column:user_id"`
	Balance   float64   `gorm:"column:balance;type:double precision;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	// has many
	Transaction []Transaction `gorm:"foreignKey:AccountID;references:ID"`
	// belongs to
	Customer *Customer `gorm:"foreignKey:UserID;references:ID"`
}

// set the table name
func (u *Account) TableName() string {
	return "accounts"
}
