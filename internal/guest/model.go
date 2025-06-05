package guest

import "time"

type Guest struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Phone       string    `gorm:"type:varchar(20);not null"`
	Relation    string    `gorm:"type:enum('family','friend','relative','other');not null"`
	Side        string    `gorm:"type:enum('bride','groom');not null"`
	IsInvited   bool      `gorm:"default:false"`
	IsAttending bool      `gorm:"default:false"`
	GuestCount  int       `gorm:"default:1"`
	AttendingGuestCount int `gorm:"default:0"`
	HasResponded bool      `gorm:"default:false"`
	ResponseSource string    `gorm:"type:varchar(50);default:'website';enum('website','whatsapp','other')"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}