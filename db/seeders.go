package db

import (
	"log"

	"github.com/qahta0/movies/models"
	"gorm.io/gorm"
)

func RunSeeders(db *gorm.DB, isAllowedToRunSeeders bool) {
	if isAllowedToRunSeeders {
		seedUsers(db)
	}
}

func seedUsers(db *gorm.DB) {
	var userCount uint
	user := models.User{
		ID:       1,
		Email:    "abdullah.n.alqahtani@hotmail.com",
		Username: "qahta0",
		Password: "qahta0",
		Name:     "Abdullah Alqahtani",
	}
	db.Model(&models.User{}).Where("email = ?", user.Email).Count(&userCount)
	if userCount == 0 {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("err: %v", err)
		}
	} else {
		log.Println("User already exists, skipping seed.")
	}
}
