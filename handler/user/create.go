package user

import (
	. "github.com/yilingfeng/apiserver/handler"
	"github.com/gin-gonic/gin"
	"github.com/yilingfeng/apiserver/pkg/errno"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/yilingfeng/apiserver/util"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	
	// var r struct {
	// 	Username string `json:"username"`
	// 	Password string `json:"password"`
	// }
	var r CreateRequest
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
	
	// if r.Username == "" {
	// 	// err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
	// 	// log.Errorf(err, "Get an error")
	// 	SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
	// 	return
	// }
	
	// if errno.IsErrUserNotFound(err) {
	// 	log.Debug("err type is ErrUserNotFound")
	// }
	
	// if r.Password == "" {
	// 	// err = fmt.Errorf("password is empty")
	// 	SendResponse(c, fmt.Errorf("password is empty"), nil)
	// 	return
	// }
	if err := r.checkParam(); err != nil {
		SendResponse(c, err, nil)
		return
	}
	
	rsp := CreateResponse{
		Username: r.Username,
	}
	
	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
	SendResponse(c, nil, rsp)
}

func (r *CreateRequest) checkParam() error {
	
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}
	
	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}
	
	return nil
}
