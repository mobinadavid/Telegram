package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func GroupMemberSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample group members to seed
	groupMembers := []models.GroupMemberModel{
		{
			GroupID:   6,
			AccountID: 1,
			IsAdmin:   true,
			JoinedAt:  time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			GroupID:   7,
			AccountID: 2,
			IsAdmin:   false,
			JoinedAt:  time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			GroupID:   7,
			AccountID: 3,
			IsAdmin:   false,
			JoinedAt:  time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			GroupID:   7,
			AccountID: 4,
			IsAdmin:   true,
			JoinedAt:  time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			GroupID:   8,
			AccountID: 5,
			IsAdmin:   false,
			JoinedAt:  time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// SQL query to insert group members
	query := `INSERT INTO group_members (group_id, account_id, is_admin, joined_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5)`

	// Insert each group member
	for _, groupMember := range groupMembers {
		_, err := db.Exec(context.Background(), query, groupMember.GroupID, groupMember.AccountID, groupMember.IsAdmin, groupMember.JoinedAt, groupMember.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed group member for group ID %d, account ID %d: %v", groupMember.GroupID, groupMember.AccountID, err)
			continue
		}
	}

	log.Println("Successfully seeded group members")
}
