package repository

import (
	"errors"
	"ticket-system/internal/model"
	"ticket-system/internal/pkg/errcode"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errcode.ErrTicketNotFound
	}
	return tickets, err
}

func (r *TicketRepository) GetById(id uint) (*model.Ticket, error) {
	var ticket model.Ticket
	err := r.db.First(&ticket, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errcode.ErrTicketNotFound
	}
	return &ticket, err
}

func (r *TicketRepository) UpdateStatus(id uint, status model.TicketStatus) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var ticket model.Ticket
		// 使用乐观锁，先锁定记录 再更新 状态和版本号
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&ticket, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errcode.ErrTicketNotFound
			}
			return err
		}

		res := tx.Model(&model.Ticket{}).
			Where("id = ? AND version = ?", id, ticket.Version).
			Updates(map[string]interface{}{"status": status, "version": ticket.Version + 1})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return errcode.ErrConcurrentUpdate
		}
		return nil
	})
}

func (r *TicketRepository) Assign(id uint, userId uint) error {
	result := r.db.Model(&model.Ticket{}).Where("id = ?", id).Update("assigned_to", userId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errcode.ErrTicketNotFound
	}
	return nil
}
