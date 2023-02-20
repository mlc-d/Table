package internal

import (
	"context"
	"sync"

	"gitlab.com/mlc-d/table/db"
)

type repo struct{}

var (
	r         *repo
	singleton sync.Once
)

type Repo interface {
	Insert(context.Context, *User) (int64, error)
	GetPasswordByName(context.Context, *User) (*User, error)
}

func GetRepo() Repo {
	singleton.Do(func() {
		r = &repo{}
	})
	return r
}

func (repo) Insert(ctx context.Context, user *User) (int64, error) {
	return db.Insert(ctx, user)
}

func (repo) GetPasswordByName(ctx context.Context, user *User) (*User, error) {
	return db.Select(ctx,
		user,
		"user.name",
		user.Name,
		"id", "name", "password", "role_id",
	)
}
