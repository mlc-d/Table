package main

import (
	"log"

	"gitlab.com/mlc-d/table/web/api"
)

func main() {
	server := api.GetServer()
	log.Fatalln(server.Start())
}
