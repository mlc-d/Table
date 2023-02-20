package internal

import (
	"context"
	"sync"

	"gitlab.com/mlc-d/table/db"
)

type repo struct {}

var (
	r         *repo
	singleton sync.Once
)

type Repo interface {
	Insert(context.Context, *Topic) (int64, error)
}

func GetRepo() Repo {
	singleton.Do(func() {
		r = &repo{}
	})
	return r
}

func (repo) Insert(ctx context.Context, topic *Topic) (int64, error) {
	return db.Insert(ctx, topic)
}
