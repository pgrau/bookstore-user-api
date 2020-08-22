package app

import (
	"github.com/pgrau/bookstore-user-api/controller/ping"
	"github.com/pgrau/bookstore-user-api/controller/user"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
}
