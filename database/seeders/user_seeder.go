package seeders

import (
	"log"
	"safrenz-go-boilerplate/internal/model"

	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	users := []model.User{
		{
			Name:     "Admin Safrenz",
			Email:    "admin@safrenz.com",
			Password: "password",
			Role:     "admin",
		},
		{
			Name:     "Intern Member",
			Email:    "intern@safrenz.com",
			Password: "password",
			Role:     "intern",
		},
	}

	for _, user := range users {
		err := db.Where(model.User{Email: user.Email}).FirstOrCreate(&user).Error
		if err != nil {
			log.Printf("Gagal seeding user %s: %v", user.Email, err)
		}
	}
	log.Println("Seeder User berhasil dijalankan!")
}
