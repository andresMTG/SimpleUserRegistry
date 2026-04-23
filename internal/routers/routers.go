package routers

import (
	"andresMTG/SimpleUserRegistry/internal/config"
	"andresMTG/SimpleUserRegistry/internal/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)



func Router(cfg *config.AppConfig) {
	router := gin.Default()

	addr := fmt.Sprintf(":%s", cfg.Port)

	v1 := router.Group("api/v1")

	v1.GET("/health", controllers.Health)

	router.Run(addr)
}
