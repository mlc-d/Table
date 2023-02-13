package main

import (
	"gitlab.com/mlc-d/table/db"
	"gitlab.com/mlc-d/table/web/api"
)

func main() {
	// test if the database opens
	_ = db.GetDB()
	server := api.GetServer()
	server.Start()
}
