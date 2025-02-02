package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func AccountSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	accounts := []models.Account{
		{Username: "Ali", Password: "hashed_password1", Mobile: "09122590238", Email: "ali@example.com", IsPremium: false},
		{Username: "Hossein", Password: "hashed_password2", Mobile: "09126844764", Email: "hossein@example.com", IsPremium: false},
		{Username: "Mohammad", Password: "hashed_password3", Mobile: "09302598945", Email: "mohammad@example.com", IsPremium: false},
		{Username: "Mobina", Password: "hashed_password4", Mobile: "09388835773", Email: "mobina@example.com", IsPremium: true},
		{Username: "Tina", Password: "hashed_password5", Mobile: "09389606266", Email: "tina@example.com", IsPremium: false},
		{Username: "Sarina", Password: "hashed_password6", Mobile: "09362589749", Email: "sarina@example.com", IsPremium: true},
	}

	query := `INSERT INTO accounts (username, password, mobile, email,is_premium, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $7,$6) ;`

	for _, account := range accounts {
		_, err := db.Exec(context.Background(), query, account.Username, account.Password, account.Mobile, account.Email, account.IsPremium, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed account %s: %v", account.Username, err)
		}
	}
	log.Printf("Successfully seeded account ")

}
