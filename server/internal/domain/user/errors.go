package user

import "errors"

var (
	ErrNotFound            = errors.New("user: not found")
	ErrInvalidCredentials  = errors.New("user: invalid credentials")
	ErrInvalidUsername     = errors.New("user: invalid username")
	ErrInvalidPasswordHash = errors.New("user: invalid password hash")
	ErrAlreadyExists       = errors.New("user: already exists")
)
