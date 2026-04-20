package security

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// PasswordGenerator génère des mots de passe aléatoires humain-lisibles.
// Alphabet réduit volontairement : exclusion de caractères ambigus (0/O, 1/l/I)
// pour qu'un humain puisse recopier sans erreur depuis les logs.
type PasswordGenerator struct{}

func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{}
}

const alphabet = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

func (g *PasswordGenerator) Generate(length int) (string, error) {
	if length < 16 {
		return "", errors.New("password: length must be >= 16")
	}
	if length > 128 {
		return "", errors.New("password: length must be <= 128")
	}
	out := make([]byte, length)
	max := big.NewInt(int64(len(alphabet)))
	for i := range out {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		out[i] = alphabet[n.Int64()]
	}
	return string(out), nil
}
