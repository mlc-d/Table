package dto

import "time"

type Comment struct {
	ID        int64      `json:"id,omitempty"`
	UserID    int64      `json:"user_id,omitempty"`
	ThreadID  int64      `json:"thread_id,omitempty"`
	Tag       string     `json:"tag,omitempty"`
	Body      string     `json:"body,omitempty"`
	Media     *Media     `json:"media,omitempty"`
	IsOP      bool       `json:"is_op,omitempty"`
	Color     uint8      `json:"color,omitempty"`
	UniqueID  string     `json:"unique_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Comments []*Comment
