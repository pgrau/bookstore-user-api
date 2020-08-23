package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pgrau/bookstore-user-api/domain/user"
	"github.com/pgrau/bookstore-user-api/service"
	"github.com/pgrau/bookstore-user-api/lib/error"
	"net/http"
	"strconv"
)

func Get(c *gin.Context)  {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := service.GetUser(userId)
	if (getErr != nil) {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Create(c *gin.Context)  {
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

func Update(c *gin.Context)  {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := error.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, saveErr := service.UpdateUser(isPartial, user)
	if (saveErr != nil) {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context)  {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	deleteErr := service.DeleteUser(userId)
	if (deleteErr != nil) {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}

func getUserId(userIdParam string)(int64, *error.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, error.BadRequest("user id should be a number")
	}

	return userId, nil
}