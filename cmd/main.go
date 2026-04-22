package main

import (
	"andresMTG/SimpleUserRegistry/internal/config"
	"andresMTG/SimpleUserRegistry/internal/routers"
)

func main() {

	cfg := config.Load()

	routers.Router(cfg)
}
