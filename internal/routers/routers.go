package routers

import (
	"andresMTG/SimpleUserRegistry/internal/controllers"
	"andresMTG/SimpleUserRegistry/internal/config"
	"github.com/gin-gonic/gin"
)


func Router(cfg *config.AppConfig) {
	router := gin.Default()
	
	router.GET("/health", controllers.Health)

	router.Run(cfg.Addrs)
}