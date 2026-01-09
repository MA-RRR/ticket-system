package service

import (
	"ticket-system/internal/model"
	"ticket-system/internal/pkg/errcode"
	"ticket-system/internal/repository"
)

type TicketService struct {
	ticketRepository *repository.TicketRepository
	auditService     *AuditService
}

func NewTicketService(ticketRepository *repository.TicketRepository) *TicketService {
	return &TicketService{ticketRepository: ticketRepository}
}

func (s *TicketService) SetAuditService(auditService *AuditService) {
	s.auditService = auditService
}

func (s *TicketService) Create(title, desc string) (*model.Ticket, error) {
	ticket := &model.Ticket{Title: title, Description: desc, Status: model.StatusNew, Version: 1}
	err := s.ticketRepository.Create(ticket)
	return ticket, err
}

func (s *TicketService) List() ([]model.Ticket, error) {
	return s.ticketRepository.List()
}

func (s *TicketService) GetById(id uint) (*model.Ticket, error) {
	return s.ticketRepository.GetById(id)
}

func (s *TicketService) UpdateStatus(id uint, newStatus model.TicketStatus, userID uint) error {
	if !IsVaildStatus(newStatus) {
		return errcode.ErrInvalidStatusTransfer
	}

	ticket, err := s.ticketRepository.GetById(id)
	if err != nil {
		return err
	}
	if !CanTranfer(ticket.Status, newStatus) {
		return errcode.ErrInvalidStatusTransfer
	}
	err = s.ticketRepository.UpdateStatus(id, newStatus)
	if err != nil {
		return err
	}
	// Log audit
	if s.auditService != nil {
		s.auditService.LogStatusUpdate(id, userID, ticket.Status, newStatus)
	}
	return nil
}

func (s *TicketService) Assign(id uint, userId uint, operatorID uint) error {
	if userId == 0 {
		return errcode.ErrInvalidParam
	}
	ticket, err := s.ticketRepository.GetById(id)
	if err != nil {
		return err
	}
	if ticket == nil {
		return errcode.ErrTicketNotFound
	}
	err = s.ticketRepository.Assign(id, userId)
	if err != nil {
		return err
	}
	// Log audit
	if s.auditService != nil {
		s.auditService.LogAssign(id, operatorID, userId)
	}
	return nil
}

func IsVaildStatus(s model.TicketStatus) bool {
	switch s {
	case model.StatusNew, model.StatusProcessing, model.StatusClosed:
		return true

	default:
		return false
	}
}
