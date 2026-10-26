package main

import (
	"log"

	"github.com/amirmalekian/Mashinchi/src/api"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/data/cache"
	"github.com/amirmalekian/Mashinchi/src/data/db"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDB(cfg)
	defer db.CloseDB()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
