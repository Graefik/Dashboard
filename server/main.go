package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Utilisateur struct {
	ID             uint   `gorm:"primaryKey"`
	Email          string `gorm:"uniqueIndex;size:255;not null"`
	MotDePasse     string `gorm:"size:255;not null"`
	NomUtilisateur string `gorm:"uniqueIndex;size:100;not null"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!"})
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "dashboard"),
		getEnv("DB_PASSWORD", "dashboard_pwd"),
		getEnv("DB_HOST", "mysql"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "dashboard_dev"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connexion DB échouée: %v", err)
	}

	if err := db.AutoMigrate(&Utilisateur{}); err != nil {
		log.Fatalf("migration échouée: %v", err)
	}
	var user Utilisateur
	result := db.Where("nom_utilisateur = ?", "admin").First(&user)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		pass, err2 := GeneratePassword(24)
		if err2 != nil {
			log.Println("Error when creating password")
		}
		user := Utilisateur{Email: "", MotDePasse: pass, NomUtilisateur: "admin"}
		log.Println("Mot de passe" + pass)
		db.Create(&user)
	}
	log.Println("Mot de passe " + user.MotDePasse)
	log.Println("DB prête.")
	r.Run(":8080")
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

func GeneratePassword(n int) (string, error) {
	if n > 24 {
		n = 24
	}

	pass := make([]byte, n)
	for i := range pass {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		pass[i] = letters[num.Int64()]
	}

	return string(pass), nil
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
