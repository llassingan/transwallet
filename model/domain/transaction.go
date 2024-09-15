package domain

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey;column:id;autoIncrement"`
	AccountID uint      `gorm:"column:account"`
	Amount    float64   `gorm:"type:double precision;column:amount"`
	Type      string    `gorm:"column:type"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	// belongs to
	AccountInfo Account `gorm:"foreignKey:AccountID;references:ID"`
}

// set the table name
func (u *Transaction) TableName() string {
	return "transactions"
}
