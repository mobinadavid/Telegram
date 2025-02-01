package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
)

func AccountRoleSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample account-role assignments
	accountRoles := []struct {
		AccountID int64
		RoleID    int64
	}{
		{AccountID: 1, RoleID: 1}, // Account 1 gets the 'Admin' role (assuming RoleID 1 is 'Admin')
		{AccountID: 2, RoleID: 2}, // Account 2 gets the 'User' role (assuming RoleID 2 is 'User')
		{AccountID: 3, RoleID: 3}, // Account 3 gets the 'Moderator' role
		{AccountID: 4, RoleID: 4}, // Account 4 gets the 'premium' role
		{AccountID: 5, RoleID: 2}, // Account 5 gets the 'SuperAdmin' role
	}

	// SQL query to insert account-role mappings
	query := `INSERT INTO account_role (account_id, role_id) 
	          VALUES ($1, $2)`

	// Insert each account-role mapping
	for _, accountRole := range accountRoles {
		_, err := db.Exec(context.Background(), query, accountRole.AccountID, accountRole.RoleID)
		if err != nil {
			log.Printf("Failed to seed account-role mapping for account %d and role %d: %v", accountRole.AccountID, accountRole.RoleID, err)
		}
	}

	log.Println("Successfully seeded account-role mappings")
}
