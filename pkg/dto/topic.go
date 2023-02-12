package dto

import "time"

type Topic struct {
	ID             int64      `json:"id,omitempty"`
	ShortName      string     `json:"short_name,omitempty"`
	Name           string     `json:"name,omitempty"`
	Media          *Media     `json:"media,omitempty"`
	Threads        Threads    `json:"threads,omitempty"`
	IsNSFW         bool       `json:"is_nsfw,omitempty"`
	MaximumThreads uint8      `json:"maximum_threads,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}
