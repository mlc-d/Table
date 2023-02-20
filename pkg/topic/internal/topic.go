package internal

import (
	"time"

	"github.com/uptrace/bun"
)

type Topic struct {
	bun.BaseModel
	ID             int64      `bun:"id,pk,autoincrement"`
	ShortName      string     `bun:"shortName,notnull,unique"`
	Name           string     `bun:"name,notnull,unique"`
	IsNSFW         bool       `bun:"isNSFW,notnull,default:false"`
	MaximumThreads uint8      `bun:"maximumThreads,notnull,default:64"`
	CreatedAt      *time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt      *time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
