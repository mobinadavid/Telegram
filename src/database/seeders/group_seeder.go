package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func GroupSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample groups to seed
	groups := []models.GroupModel{
		{
			Name:         "Tech Enthusiasts",
			Description:  "A group for tech lovers to share the latest trends.",
			OwnerID:      1,
			MembersCount: 10002,
			InviteLink:   "https://example.com/invite1",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "Travelers Unite",
			Description:  "A group for travel enthusiasts to discuss destinations.",
			OwnerID:      2,
			MembersCount: 111000,
			InviteLink:   "https://example.com/invite2",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "Music Lovers",
			Description:  "A group for those passionate about music.",
			OwnerID:      3,
			MembersCount: 0,
			InviteLink:   "https://example.com/invite3",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	// SQL query to insert groups
	query := `INSERT INTO groups (name, description, owner_id, members_count, invite_link, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	// Insert each group
	for _, group := range groups {
		var groupID int64
		err := db.QueryRow(context.Background(), query, group.Name, group.Description, group.OwnerID, group.MembersCount, group.InviteLink, group.CreatedAt, group.UpdatedAt).Scan(&groupID)
		if err != nil {
			log.Printf("Failed to seed group: %s - %v", group.Name, err)
			continue
		}
	}

	log.Println("Successfully seeded groups")
}
