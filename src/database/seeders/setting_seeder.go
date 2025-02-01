package seeders

//
//import (
//	"context"
//	"log"
//	"telegram/src/database/pgx"
//	"time"
//)
//
//func SettingSeeder(accountID int64) (int64, error) {
////	db := pgx.GetInstance()
////	if db == nil {
////		log.Println("Database connection is nil")
////		return 0, nil
////	}
////
////	var settingID int64
////	query := `INSERT INTO setting (owner_id, language, created_at, updated_at)
////	          VALUES ($1, 'en', $2, $3) RETURNING id;`
////	err := db.QueryRow(context.Background(), query, accountID, time.Now(), time.Now()).Scan(&settingID)
////	if err != nil {
////		log.Printf("Failed to seed setting for account %d: %v", accountID, err)
////		return 0, err
////	}
////	return settingID, nil
////}
////
////func PrivacyAndSecuritySeeder(settingID int64) {
////	db := pgx.GetInstance()
////	if db == nil {
////		log.Println("Database connection is nil")
////		return
////	}
////
////	query := `INSERT INTO privacy_and_security (setting_id, last_seen_visibility, profile_visibility,
////	           phone_number_visibility, bio_visibility, two_step_verification, passcode_lock, created_at, updated_at)
////	          VALUES ($1, 'everyone', 'everyone', 'everyone', 'everyone', false, false, $2, $3);`
////
////	_, err := db.Exec(context.Background(), query, settingID, time.Now(), time.Now())
////	if err != nil {
////		log.Printf("Failed to seed privacy_and_security for setting %d: %v", settingID, err)
////	}
//}
//
//func NotificationAndSoundSeeder(settingID int64) {
//	db := pgx.GetInstance()
//	if db == nil {
//		log.Println("Database connection is nil")
//		return
//	}
//
//	query := `INSERT INTO notification_and_sound (setting_id, vibrate, ringtone, group_notif, channel_notif, created_at, updated_at)
//	          VALUES ($1, 'default', 'default', true, true, $2, $3);`
//
//	_, err := db.Exec(context.Background(), query, settingID, time.Now(), time.Now())
//	if err != nil {
//		log.Printf("Failed to seed notification_and_sound for setting %d: %v", settingID, err)
//	}
//}
