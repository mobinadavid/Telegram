package models

type ContactModel struct {
	AccountID  int64 `json:"account_id"`
	ContactID  int64 `json:"contact_id"`
	IsFavorite bool  `json:"is_favorite"`
	IsBlocked  bool  `json:"is_blocked"`
}
