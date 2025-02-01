package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func ProfilePictureSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	profilePictures := []models.ProfilePicture{
		{PictureURL: "https://example.com/profiles/user1.jpg", OwnerID: 1, OwnerType: "account", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PictureURL: "https://example.com/profiles/group1.jpg", OwnerID: 6, OwnerType: "group", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PictureURL: "https://example.com/profiles/channel1.jpg", OwnerID: 3, OwnerType: "channel", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PictureURL: "https://example.com/profiles/user2.jpg", OwnerID: 2, OwnerType: "account", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PictureURL: "https://example.com/profiles/group2.jpg", OwnerID: 7, OwnerType: "group", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	query := `INSERT INTO profile_pictures (picture_url, owner_id, owner_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (owner_id, owner_type) DO NOTHING;`

	for _, pp := range profilePictures {
		_, err := db.Exec(context.Background(), query, pp.PictureURL, pp.OwnerID, pp.OwnerType, pp.CreatedAt, pp.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed profile picture %s: %v", pp.PictureURL, err)
		}
	}
	log.Println("Successfully seeded profile pictures")
}
