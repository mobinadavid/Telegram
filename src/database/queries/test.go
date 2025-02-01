package queries

import (
	"context"
	"fmt"
	"log"
	"telegram/src/database/pgx"
)

func getPremiumAccounts() {
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT id, username, email FROM accounts WHERE is_premium = true")
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var username, email string
		if err := rows.Scan(&id, &username, &email); err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", id, username, email)
	}

	if rows.Err() != nil {
		log.Fatal("Error while iterating rows:", rows.Err())
	}
}
