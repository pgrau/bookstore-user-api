package app

import (
	"github.com/pgrau/bookstore-user-api/controller/ping"
	"github.com/pgrau/bookstore-user-api/controller/user"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/user", user.CreateUser)
	router.GET("/user/:user_id", user.GetUser)
}
