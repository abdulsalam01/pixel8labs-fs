package entity

import "time"

type Repository struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Fullname    string    `json:"full_name"`
	IsPrivate   bool      `json:"private"`
	Description string    `json:"description"`
	Topics      []string  `json:"topics"`
	UpdatedAt   time.Time `json:"updated_at"`
	Language    string    `json:"language"`
	Url         string    `json:"html_url"`
}
