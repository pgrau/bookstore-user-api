package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pgrau/bookstore-user-api/domain/user"
	"github.com/pgrau/bookstore-user-api/service"
	"github.com/pgrau/bookstore-user-api/lib/error"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := error.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := service.GetUser(userId)
	if (getErr != nil) {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
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
