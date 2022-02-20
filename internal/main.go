package main

import (
	"fmt"
	"log"
	sql2 "main/internal/database/sql"
	server2 "main/internal/server"
	_ "net/http/pprof"
)

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	sql2.Connect()
	server2.Init()

	_, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
