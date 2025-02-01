package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func StoryViewSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample story view data
	storyViews := []models.StoryView{
		{ViewerID: 1, StoryID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ViewerID: 2, StoryID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ViewerID: 3, StoryID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ViewerID: 4, StoryID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ViewerID: 5, StoryID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ViewerID: 6, StoryID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// SQL query to insert data into story_views table
	query := `INSERT INTO story_views (viewer_id, story_id, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4);`

	// Loop through the storyViews and insert each one into the database
	for _, view := range storyViews {
		_, err := db.Exec(context.Background(), query, view.ViewerID, view.StoryID, view.CreatedAt, view.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed story view for story ID %d: %v", view.StoryID, err)
		}
	}

	log.Println("Successfully seeded story views")
}
