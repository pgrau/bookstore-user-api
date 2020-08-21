package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pgrau/bookstore-user-api/domain/user"
	"github.com/pgrau/bookstore-user-api/service"
	"github.com/pgrau/bookstore-user-api/lib/error"
	"net/http"
)

func GetUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context)  {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := error.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := service.CreateUser(user)
	if (saveErr != nil) {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
