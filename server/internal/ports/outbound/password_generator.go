package outbound

// PasswordGenerator génère un mot de passe aléatoire humain-lisible
// (utilisé au premier boot pour créer l'admin Jenkins-style).
type PasswordGenerator interface {
	Generate(length int) (string, error)
}
