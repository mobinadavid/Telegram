package models

import "time"

type Storage struct {
	ID         int64      `json:"id"`
	FileName   string     `json:"file_name"`
	FilePath   string     `json:"file_path"`
	FileSize   int64      `json:"file_size"`
	FileType   string     `json:"file_type"`
	MimeType   string     `json:"mime_type"`
	UploaderID int64      `json:"uploader_id"`
	MessageID  *int64     `json:"message_id,omitempty"`
	ChatID     int64      `json:"chat_id"`
	UploadedAt time.Time  `json:"uploaded_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
