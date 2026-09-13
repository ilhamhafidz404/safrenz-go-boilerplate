package seeders

import (
	"log"

	"gorm.io/gorm"
)

// RunAll melempar perintah eksekusi ke setiap seeder individual
func RunAll(db *gorm.DB) {
	log.Println("Memulai proses database seeding...")

	UserSeeder(db)
	// Jika ada seeder lain nanti:
	// ProductSeeder(db)
	// RoleSeeder(db)

	log.Println("Proses database seeding selesai!")
}
