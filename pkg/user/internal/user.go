package internal

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel
	ID        int64      `bun:"id,pk,autoincrement"`
	Name      string     `bun:"name,notnull,unique"`
	Password  string     `bun:"password,notnull"`
	RoleID    uint8      `bun:"role_id,notnull"`
	IsBanned  bool       `bun:"is_banned,default:false"`
	CreatedAt *time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
