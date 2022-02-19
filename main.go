package main

import (
	"fmt"
	"log"
	"main/database/sql"
	"main/server"
	_ "net/http/pprof"
)

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	sql.Connect()
	server.Init()

	_, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
