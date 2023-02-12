package db

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

var (
	db        *bun.DB
	singleton sync.Once
	dsn       = os.Getenv("DSN")
)

func GetDB() *bun.DB {
	singleton.Do(func() {
		sqldb, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		db = bun.NewDB(sqldb, mysqldialect.New())
	})
	return db
}
