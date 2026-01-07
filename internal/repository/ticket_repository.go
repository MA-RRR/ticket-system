package repository

import (
	"ticket-system/internal/model"

	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) Create(ticket *model.Ticket) error {
	return r.db.Create(ticket).Error
}
func (r *TicketRepository) List() ([]model.Ticket, error) {
	var tickets []model.Ticket
	err := r.db.Find(&tickets).Error
	return tickets, err
}

func (r *TicketRepository) GetById(id uint) (*model.Ticket, error) {
	var ticket model.Ticket
	err := r.db.First(&ticket, id).Error
	return &ticket, err
}

func (r *TicketRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&model.Ticket{}).
		Where("id = ?", id).Update("status", status).Error
}
