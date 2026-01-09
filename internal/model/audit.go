package model

import "time"

// AuditAction 审计操作类型
type AuditAction string

const (
	AuditActionStatusUpdate AuditAction = "STATUS_UPDATE"
	AuditActionAssign       AuditAction = "ASSIGN"
)

// AuditLog 审计日志
type AuditLog struct {
	ID        uint        `gorm:"primaryKey"`
	TicketID  uint        `gorm:"index;not null"`
	UserID    uint        `gorm:"not null"`
	Action    AuditAction `gorm:"size:32;not null"`
	OldValue  string      `gorm:"type:text"`
	NewValue  string      `gorm:"type:text"`
	Remark    string      `gorm:"type:text"`
	CreatedAt time.Time
}
