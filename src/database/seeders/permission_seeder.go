package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func PermissionSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	// Sample permissions to seed
	permissions := []models.PermissionModel{
		{
			Name:      "view_dashboard",
			Title:     `{"en": "View Dashboard", "fa": "مشاهده داشبورد"}`,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "manage_users",
			Title:     `{"en": "Manage Users", "fa": "مدیریت کاربران"}`,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "edit_profile",
			Title:     `{"en": "Edit Profile", "fa": "ویرایش پروفایل"}`,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "delete_account",
			Title:     `{"en": "Delete Account", "fa": "حذف حساب کاربری"}`,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "manage_roles",
			Title:     `{"en": "Manage Roles", "fa": "مدیریت نقش‌ها"}`,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "create_story",
			Title:     `{"en": "Create Story", "fa": "ایجاد استوری"}`,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// SQL query to insert permissions
	query := `INSERT INTO permissions (name, title, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4)`

	// Insert each permission
	for _, permission := range permissions {
		_, err := db.Exec(context.Background(), query, permission.Name, permission.Title, permission.CreatedAt, permission.UpdatedAt)
		if err != nil {
			log.Printf("Failed to seed permission: %s - %v", permission.Name, err)
		}
	}

	log.Println("Successfully seeded permissions")
}
