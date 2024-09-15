package domain

import "time"

type Customer struct {
	ID            uint      `gorm:"primaryKey;column:id;autoIncrement"`
	Name          string    `gorm:"column:name"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	// has one 
	AccountNumber Account   `gorm:"foreignKey:UserID;references:ID"`
}

// set the table name
func (u *Customer) TableName() string {
	return "customers"
}
