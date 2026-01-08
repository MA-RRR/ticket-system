package errcode

import "errors"

var (
	ErrInvalidStatusTransfer = errors.New("invalid status transfer")
	ErrTicketNotFound        = errors.New("ticket not found")
	ErrConcurrentUpdate      = errors.New("concurrent update")
)
