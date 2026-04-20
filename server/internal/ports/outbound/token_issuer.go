package outbound

import "time"

type TokenIssuer interface {
	Issue(userID uint, username string) (token string, expiresAt time.Time, err error)
}
