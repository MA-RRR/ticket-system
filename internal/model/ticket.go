package model

import "time"

// 工单模型
type Ticket struct {
	ID          uint   `grom:"primary key"`
	Title       string `grom:"size:255;not null"`
	Description string `grom:"type:text"`
	Status      string `grom:"size:232;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
