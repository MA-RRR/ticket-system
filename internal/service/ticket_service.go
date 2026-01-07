package service

import (
	"ticket-system/internal/model"
	"ticket-system/internal/repository"
)

type TicketService struct {
	ticketRepository *repository.TicketRepository
}

func NewTicketService(ticketRepository *repository.TicketRepository) *TicketService {
	return &TicketService{ticketRepository: ticketRepository}
}

func (s *TicketService) Create(title, desc string) (*model.Ticket, error) {
	ticket := &model.Ticket{Title: title, Description: desc, Status: "New"}
	err := s.ticketRepository.Create(ticket)
	return ticket, err
}

func (s *TicketService) List() ([]model.Ticket, error) {
	return s.ticketRepository.List()
}

func (s *TicketService) GetById(id uint) (*model.Ticket, error) {
	return s.ticketRepository.GetById(id)
}

func (s *TicketService) UpdateStatus(id uint, status string) error {
	return s.ticketRepository.UpdateStatus(id, status)
}
