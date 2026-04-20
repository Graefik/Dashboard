package traefik

import "errors"

var (
	ErrUnavailable   = errors.New("traefik: instance unreachable")
	ErrUnauthorized  = errors.New("traefik: invalid credentials")
	ErrNotFound      = errors.New("traefik: resource not found")
	ErrMalformed     = errors.New("traefik: malformed response")
)
