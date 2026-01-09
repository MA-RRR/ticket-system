package service

import (
	"fmt"
	"ticket-system/internal/model"
	"ticket-system/internal/repository"
)

type AuditService struct {
	auditRepository *repository.AuditRepository
}

func NewAuditService(auditRepository *repository.AuditRepository) *AuditService {
	return &AuditService{auditRepository: auditRepository}
}

func (s *AuditService) LogStatusUpdate(ticketID, userID uint, oldStatus, newStatus model.TicketStatus) error {
	audit := &model.AuditLog{
		TicketID: ticketID,
		UserID:   userID,
		Action:   model.AuditActionStatusUpdate,
		OldValue: string(oldStatus),
		NewValue: string(newStatus),
		Remark:   fmt.Sprintf("Status changed from %s to %s", oldStatus, newStatus),
	}
	return s.auditRepository.Create(audit)
}

func (s *AuditService) LogAssign(ticketID, userID, assignedTo uint) error {
	audit := &model.AuditLog{
		TicketID: ticketID,
		UserID:   userID,
		Action:   model.AuditActionAssign,
		NewValue: fmt.Sprintf("%d", assignedTo),
		Remark:   fmt.Sprintf("Assigned to user %d", assignedTo),
	}
	return s.auditRepository.Create(audit)
}

func (s *AuditService) ListByTicketID(ticketID uint) ([]model.AuditLog, error) {
	return s.auditRepository.ListByTicketID(ticketID)
}
