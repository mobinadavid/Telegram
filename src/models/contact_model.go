package models

// ContactModel represents the structure of the contact table in the database
type ContactModel struct {
	AccountID  int64 `json:"account_id"`
	ContactID  int64 `json:"contact_id"`
	IsFavorite bool  `json:"is_favorite"`
	IsBlocked  bool  `json:"is_blocked"`
}
