package model

type TicketStatus string

const (
	StatusNew        TicketStatus = "NEW"
	StatusProcessing TicketStatus = "PROCESSING"
	StatusClosed     TicketStatus = "CLOSED"
)
