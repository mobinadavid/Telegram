package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
)

func RolePermissionSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Example role to permission associations
	rolePermissions := []struct {
		RoleID       int64
		PermissionID int64
	}{
		// Admin - All permissions
		{RoleID: 1, PermissionID: 1}, // view_dashboard
		{RoleID: 1, PermissionID: 2}, // manage_users
		{RoleID: 1, PermissionID: 3}, // edit_profile
		{RoleID: 1, PermissionID: 4}, // delete_account
		{RoleID: 1, PermissionID: 5}, // manage_roles
		{RoleID: 1, PermissionID: 6}, // create_story

		// User - Basic permissions
		{RoleID: 2, PermissionID: 1}, // view_dashboard
		{RoleID: 2, PermissionID: 3}, // edit_profile

		// Premium - Additional permissions
		{RoleID: 4, PermissionID: 1}, // view_dashboard
		{RoleID: 4, PermissionID: 3}, // edit_profile
		{RoleID: 4, PermissionID: 6}, // create_story
	}

	// SQL query to insert role permissions
	query := `INSERT INTO permission_role (role_id, permission_id)
	          VALUES ($1, $2)`

	// Insert each role permission mapping
	for _, rolePermission := range rolePermissions {
		_, err := db.Exec(context.Background(), query, rolePermission.RoleID, rolePermission.PermissionID)
		if err != nil {
			log.Printf("Failed to assign permission to role: %v - %v", rolePermission.RoleID, err)
		}
	}

	log.Println("Successfully assigned permissions to roles")
}
