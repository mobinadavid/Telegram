package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func FAQSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	faqs := []models.FaqModel{
		{Question: "How to reset my password?", Answer: "You can reset your password by going to the settings page and selecting 'Reset Password'."},
		{Question: "How to change my email?", Answer: "Go to account settings and update your email under the 'Email' section."},
		{Question: "How to enable two-step verification?", Answer: "You can enable two-step verification in the privacy and security settings."},
		{Question: "How to delete my account?", Answer: "Please contact support to permanently delete your account."},
		{Question: "How to change my notification settings?", Answer: "Navigate to the notification settings and customize your preferences."},
	}

	query := `INSERT INTO frequently_asked_questions (question, answer, is_active, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5);`

	for _, faq := range faqs {
		_, err := db.Exec(context.Background(), query, faq.Question, faq.Answer, true, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed FAQ: %s - %v", faq.Question, err)
		}
	}

	log.Println("Successfully seeded FAQs")
}
