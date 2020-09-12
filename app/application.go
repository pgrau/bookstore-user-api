package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pgrau/bookstore-user-api/lib/logger"
)

var(
	router = gin.Default()
)

func StartApplication()  {
	MapUrls()

	logger.Info("Start the application")

	router.Run(":8081")
}



