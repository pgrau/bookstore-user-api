package app

import (
	"github.com/pgrau/bookstore-user-api/controller/ping"
	"github.com/pgrau/bookstore-user-api/controller/user"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", user.Create)
	router.GET("/users/:user_id", user.Get)
	router.PUT("/users/:user_id", user.Update)
	router.PATCH("/users/:user_id", user.Update)
	router.DELETE("/users/:user_id", user.Delete)

	router.GET("/internal/users/search", user.Search)
	router.POST("/users/login", user.Login)
}
