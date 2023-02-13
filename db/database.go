package db

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type DB struct {
	database *bun.DB
}

var (
	db        *DB
	singleton sync.Once
	dsn       = os.Getenv("DSN")
)

func GetDB() *DB {
	singleton.Do(func() {
		sqlDatabase, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		db = &DB{database: bun.NewDB(sqlDatabase, mysqldialect.New())}
		db.database.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
	})
	return db
}
