package repository

import (
	"ticket-system/internal/model"

	"gorm.io/gorm"
)

type AuditRepository struct {
	db *gorm.DB
}

func NewAuditRepository(db *gorm.DB) *AuditRepository {
	return &AuditRepository{db: db}
}

func (r *AuditRepository) Create(audit *model.AuditLog) error {
	return r.db.Create(audit).Error
}

func (r *AuditRepository) ListByTicketID(ticketID uint) ([]model.AuditLog, error) {
	var logs []model.AuditLog
	err := r.db.Where("ticket_id = ?", ticketID).Order("created_at DESC").Find(&logs).Error
	return logs, err
}
