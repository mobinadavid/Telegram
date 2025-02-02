package seeders

import (
	"context"
	"log"
	"telegram/src/database/pgx"
	"telegram/src/models"
	"time"
)

func StorageSeeder() {
	db := pgx.GetInstance()
	if db == nil {
		log.Println("Database connection is nil")
		return
	}

	messageID1 := int64(1)
	messageID2 := int64(2)

	storages := []models.Storage{
		{FileName: "image1.jpg", FilePath: "/uploads/images/image1.jpg", FileSize: 2048, FileType: "image", MimeType: "image/jpeg", UploaderID: 3, MessageID: &messageID1, ChatID: 10},
		{FileName: "video1.mp4", FilePath: "/uploads/videos/video1.mp4", FileSize: 10485760, FileType: "video", MimeType: "video/mp4", UploaderID: 2, MessageID: &messageID2, ChatID: 7},
		{FileName: "audio1.mp3", FilePath: "/uploads/audio/audio1.mp3", FileSize: 5120000, FileType: "audio", MimeType: "audio/mpeg", UploaderID: 1, MessageID: &messageID1, ChatID: 8},
		{FileName: "document1.pdf", FilePath: "/uploads/documents/document1.pdf", FileSize: 10240, FileType: "document", MimeType: "application/pdf", UploaderID: 4, MessageID: &messageID2, ChatID: 5},
		{FileName: "image2.png", FilePath: "/uploads/images/image2.png", FileSize: 3072, FileType: "image", MimeType: "image/png", UploaderID: 4, MessageID: nil, ChatID: 9},
	}

	query := `INSERT INTO storage (file_name, file_path, file_size, file_type, mime_type, uploader_id, message_id, chat_id, uploaded_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	for _, storage := range storages {
		_, err := db.Exec(context.Background(), query, storage.FileName, storage.FilePath, storage.FileSize, storage.FileType, storage.MimeType, storage.UploaderID, storage.MessageID, storage.ChatID, time.Now(), time.Now())
		if err != nil {
			log.Printf("Failed to seed storage for file %s: %v", storage.FileName, err)
		}
	}

	log.Printf("Successfully seeded storage files")
}
