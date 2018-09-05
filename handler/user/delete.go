package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scaat/apiserver/handler"
	"github.com/scaat/apiserver/model"
	"github.com/scaat/apiserver/pkg/errno"
)

// Delete deletes an user by the user identifier.
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userID)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)

}
