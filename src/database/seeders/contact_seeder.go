package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
)

func ContactSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Example contacts data to seed
	contacts := []models.ContactModel{
		{AccountID: 1, ContactID: 2, IsFavorite: true, IsBlocked: false},  // Ali & Hossein
		{AccountID: 1, ContactID: 3, IsFavorite: false, IsBlocked: false}, // Ali & Mohammad
		{AccountID: 2, ContactID: 4, IsFavorite: true, IsBlocked: false},  // Hossein & Mobina
		{AccountID: 3, ContactID: 5, IsFavorite: false, IsBlocked: false}, // Mohammad & Tina
		{AccountID: 4, ContactID: 6, IsFavorite: true, IsBlocked: false},  // Mobina & Sarina
		{AccountID: 5, ContactID: 1, IsFavorite: true, IsBlocked: false},  // Tina & Ali
	}

	// SQL query to insert contacts
	query := `INSERT INTO contacts (account_id, contact_id, is_favorite, is_blocked)
	          VALUES ($1, $2, $3, $4)`

	// Insert each contact
	for _, contact := range contacts {
		_, err := db.Exec(context.Background(), query, contact.AccountID, contact.ContactID, contact.IsFavorite, contact.IsBlocked)
		if err != nil {
			log.Printf("Failed to seed contact: %v - %v %s", contact.AccountID, contact.ContactID, err.Error())
		}
	}

	log.Println("Successfully seeded contacts")
}
