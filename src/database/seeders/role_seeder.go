package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func RoleSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample roles to seed
	roles := []models.RoleModel{
		{Title: "Admin", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Title: "User", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Title: "SuperAdmin", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Title: "Premium", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// SQL query to insert roles
	query := `INSERT INTO roles (title, created_at, updated_at)
	          VALUES ($1, $2, $3)`

	// Insert each role
	for _, role := range roles {
		_, err := db.Exec(context.Background(), query, role.Title, role.CreatedAt, role.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed role: %s - %v", role.Title, err)
		}
	}

	log.Println("Successfully seeded roles")
}
