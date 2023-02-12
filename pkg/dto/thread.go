package dto

import "time"

type Thread struct {
	ID        int64      `json:"id,omitempty"`
	TopicID   int64      `json:"topic_id,omitempty"`
	UserID    int64      `json:"user_id,omitempty"`
	Hash      string     `json:"hash,omitempty"`
	Title     string     `json:"title,omitempty"`
	Body      string     `json:"body,omitempty"`
	Media     *Media     `json:"media,omitempty"`
	Comment   Comments   `json:"comment,omitempty"`
	Sticky    bool       `json:"sticky,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Threads []*Thread
