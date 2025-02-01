package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func StorySeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample story data
	stories := []models.Story{
		{OwnerID: 6, ContentType: "image", ContentURL: "/uploads/stories/image1.jpg", Caption: "My first story!", ViewsCount: 50, ExpiresAt: time.Now().Add(24 * time.Hour), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{OwnerID: 4, ContentType: "text", ContentURL: "/uploads/stories/text1.txt", Caption: "This is a text story", ViewsCount: 10, ExpiresAt: time.Now().Add(24 * time.Hour), CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// SQL query to insert data into stories table
	query := `INSERT INTO stories (owner_id, content_type, content_url, caption, views_count, expires_at, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	// Loop through the stories and insert each one into the database
	for _, story := range stories {
		_, err := db.Exec(context.Background(), query, story.OwnerID, story.ContentType, story.ContentURL, story.Caption, story.ViewsCount, story.ExpiresAt, story.CreatedAt, story.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed story for owner %d: %v", story.OwnerID, err)
		}
	}

	log.Println("Successfully seeded stories")
}
