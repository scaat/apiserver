package user

import (
	. "github.com/yilingfeng/apiserver/handler"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/yilingfeng/apiserver/model"
	"github.com/yilingfeng/apiserver/pkg/errno"
)

// Delete deletes an user by the user identifier.
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	
	if err := model.DeleteUser(uint64(userID)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	
	SendResponse(c, nil, nil)
	
}
