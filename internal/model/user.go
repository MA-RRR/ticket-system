package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:64;not null;unique"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:32;not null"`
	CreatedAt time.Time
}
