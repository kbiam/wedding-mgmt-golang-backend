package auth

import (
	"time"
)

type AdminUser struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"type:varchar(50);not null;unique"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (AdminUser) TableName() string {
	return "admin_users"
}


