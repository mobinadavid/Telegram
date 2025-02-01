package seeders

import (
	"context"
	"github.com/google/uuid"
	"log"
	"telegram/src/database/pgx"
	"time"
)

func AuthenticationSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Select some account IDs to generate authentication for
	accountIDs := []int64{1, 2, 3}

	query := `INSERT INTO authentication (owner_id, access_token, access_token_expires_at, refresh_token, refresh_token_expires_at, ip, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ;`

	for _, accountID := range accountIDs {
		accessToken := uuid.New().String()
		refreshToken := uuid.New().String()

		_, err := db.Exec(context.Background(), query, accountID, accessToken, time.Now().Add(2*time.Hour), refreshToken, time.Now().Add(2*24*time.Hour), "127.0.0.1", time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed authentication for account ID %d: %v", accountID, err)
		}
	}
	log.Printf("Successfully seeded authentication for account ")

}
