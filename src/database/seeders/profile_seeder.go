package seeders

func ProfileSeeder() {
	//db := pgx.GetInstance()
	//if db == nil {
	//	log.Println("Database connection is nil")
	//	return
	//}
	//
	//// Fetch account IDs from the database
	//rows, err := db.Query(context.Background(), "SELECT id, username FROM accounts")
	//if err != nil {
	//	log.Println("Failed to fetch account IDs:", err)
	//	return
	//}
	//defer rows.Close()
	//
	//var profiles []models.ProfileModel
	//
	//for rows.Next() {
	//	var id int64
	//	var username string
	//	if err := rows.Scan(&id, &username); err != nil {
	//		log.Println("Failed to scan account row:", err)
	//		continue
	//	}
	//
	//	profile := models.ProfileModel{
	//		OwnerID:   id,
	//		FirstName: username, // Using username as first name for example
	//		LastName:  "",       // Default last name
	//		Bio:       "",
	//	}
	//	profiles = append(profiles, profile)
	//}
	//
	//query := `INSERT INTO profile (owner_id, first_name, last_name, bio, created_at, updated_at)
	//          VALUES ($1, $2, $3, $4, $5, $6);`
	//
	//for _, profile := range profiles {
	//	_, err := db.Exec(context.Background(), query, profile.OwnerID, profile.FirstName, profile.LastName, profile.Bio, time.Now(), time.Now())
	//	if err != nil {
	//		log.Printf("Failed to seed profile for account ID %d: %v", profile.OwnerID, err)
	//	}
	//}
	//
	//log.Println("Successfully seeded profiles")
}
