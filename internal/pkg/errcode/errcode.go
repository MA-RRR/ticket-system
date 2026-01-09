package errcode

import "errors"

var (
	ErrInvalidStatusTransfer = errors.New("invalid status transfer")
	ErrTicketNotFound        = errors.New("ticket not found")
	ErrConcurrentUpdate      = errors.New("concurrent update")
	ErrUnathorized           = errors.New("unauthorized")
	ErrForbidden             = errors.New("forbidden")
	ErrInvalidParam          = errors.New("invalid parameter")
	ErrUserNotFound          = errors.New("user not found")
)
