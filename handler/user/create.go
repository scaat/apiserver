package user

import (
	. "github.com/yilingfeng/apiserver/handler"
	"github.com/gin-gonic/gin"
	"github.com/yilingfeng/apiserver/pkg/errno"
	"github.com/lexkong/log"
	"fmt"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	
	// var err error
	
	if err := c.Bind(&r); err != nil {
		// c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	
	admin := c.Param("username")
	log.Infof("URL username: %s", admin)
	
	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)
	
	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)
	
	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	
	if r.Username == "" {
		// err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		// log.Errorf(err, "Get an error")
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}
	
	// if errno.IsErrUserNotFound(err) {
	// 	log.Debug("err type is ErrUserNotFound")
	// }
	
	if r.Password == "" {
		// err = fmt.Errorf("password is empty")
		SendResponse(c, fmt.Errorf("password is empty"), nil)
		return
	}
	
	rsp := CreateResponse{
		Username: r.Username,
	}
	
	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
	SendResponse(c, nil, rsp)
}
