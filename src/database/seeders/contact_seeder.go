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

	contacts := []models.ContactModel{
		{AccountID: 1, ContactID: 2, IsFavorite: true, IsBlocked: false},  // Ali & Hossein
		{AccountID: 2, ContactID: 1, IsFavorite: true, IsBlocked: false},  // Hossein & Ali
		{AccountID: 1, ContactID: 4, IsFavorite: false, IsBlocked: false}, // Ali & Mobina
		{AccountID: 2, ContactID: 4, IsFavorite: true, IsBlocked: false},  // Hossein & Mobina
		{AccountID: 3, ContactID: 2, IsFavorite: false, IsBlocked: false}, // Mohammad & Hossein
		{AccountID: 3, ContactID: 4, IsFavorite: false, IsBlocked: false}, // Mohammad & Mobina
		{AccountID: 3, ContactID: 6, IsFavorite: false, IsBlocked: false}, // Mohammad & Sarina
		{AccountID: 4, ContactID: 6, IsFavorite: true, IsBlocked: false},  // Mobina & Sarina
		{AccountID: 5, ContactID: 1, IsFavorite: true, IsBlocked: false},  // Tina & Ali
	}

	query := `INSERT INTO contacts (account_id, contact_id, is_favorite, is_blocked)
	          VALUES ($1, $2, $3, $4)`

	for _, contact := range contacts {
		_, err := db.Exec(context.Background(), query, contact.AccountID, contact.ContactID, contact.IsFavorite, contact.IsBlocked)
		if err != nil {
			log.Printf("Failed to seed contact: %v - %v %s", contact.AccountID, contact.ContactID, err.Error())
		}
	}

	log.Println("Successfully seeded contacts")
}
