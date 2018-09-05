package user

import (
	"github.com/gin-gonic/gin"
	"github.com/scaat/apiserver/handler"
	"github.com/scaat/apiserver/model"
	"github.com/scaat/apiserver/pkg/errno"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}
