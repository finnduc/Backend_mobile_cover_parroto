package errors

import "errors"

// Domain errors — wrap at service layer
var (
	ErrNotFound       = errors.New("not found")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrInvalidInput   = errors.New("invalid input")
	ErrConflict       = errors.New("conflict")
	ErrInternalServer = errors.New("internal server error")
	ErrTokenBlacklisted = errors.New("token is blacklisted")
	ErrTokenExpired     = errors.New("token expired")
)
