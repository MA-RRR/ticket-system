package model

import "time"

// 工单模型
type Ticket struct {
	ID         uint         `gorm:"primaryKey"`
	Title      string       `gorm:"size:255;not null"`
	Description string       `gorm:"type:text"`
	Status     TicketStatus `gorm:"size:32;not null"`
	AssignedTo uint         `gorm:"default:0"`
	Version    uint         `gorm:"not null;default:1"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
