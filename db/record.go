package db

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type Record interface {
	any
}

func Insert[T Record](ctx context.Context, record T) (int64, error) {
	var id int64
	err := db.database.NewInsert().
		Model(record).
		Returning("id").
		Scan(ctx, &id)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return id, nil
}

func Select[T Record](ctx context.Context, model T, lookupColumn, searchValue string, returnFields ...string) (T, error) {
	err := db.database.NewSelect().
		Model(model).
		Column(returnFields...).
		Where("? = ?", bun.Ident(lookupColumn), searchValue).
		Scan(ctx)
	return model, err
}
