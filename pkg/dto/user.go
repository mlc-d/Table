package dto

import "time"

type User struct {
	ID        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Password  string     `json:"password,omitempty"`
	RoleID    uint8      `json:"role_id,omitempty"`
	IsBanned  bool       `json:"is_banned,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
