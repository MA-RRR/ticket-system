package service

import (
	"ticket-system/internal/model"
)

var statusFlow = map[model.TicketStatus][]model.TicketStatus{
	model.StatusNew:        {model.StatusProcessing},
	model.StatusProcessing: {model.StatusClosed},
}

func CanTranfer(from, to model.TicketStatus) bool {
	if from == to {
		return true
	}
	for _, s := range statusFlow[from] {
		if s == to {
			return true
		}
	}
	return false
}
