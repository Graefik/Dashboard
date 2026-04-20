package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// All retourne l'ensemble des migrations dans l'ordre d'exécution.
// Chaque migration a un ID daté (YYYYMMDDHHMM_...).
func All() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		m202604200001InitUsers(),
	}
}

// Run applique toutes les migrations en attente.
func Run(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, All())
	return m.Migrate()
}
