package service

import (
	"ticket-system/internal/model"
	"ticket-system/internal/pkg/errcode"
	"ticket-system/internal/repository"
)

type TicketService struct {
	ticketRepository *repository.TicketRepository
}

func NewTicketService(ticketRepository *repository.TicketRepository) *TicketService {
	return &TicketService{ticketRepository: ticketRepository}
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

func (s *TicketService) UpdateStatus(id uint, newStatus model.TicketStatus) error {
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
	return s.ticketRepository.UpdateStatus(id, newStatus)
}

func IsVaildStatus(s model.TicketStatus) bool {
	switch s {
	case model.StatusNew, model.StatusProcessing, model.StatusClosed:
		return true

	default:
		return false
	}
}
