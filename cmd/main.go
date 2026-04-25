package main

import (
	"andresMTG/SimpleUserRegistry/internal/config"
	"andresMTG/SimpleUserRegistry/internal/database"
	"andresMTG/SimpleUserRegistry/internal/routers"
	
)

func main() {

	cfg := config.Load()
	
	db := database.Connect(cfg)
	defer db.Close()

	routers.Router(cfg)
}
