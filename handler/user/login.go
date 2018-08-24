package user

import (
	. "github.com/yilingfeng/apiserver/handler"
	"github.com/gin-gonic/gin"
	"github.com/yilingfeng/apiserver/model"
	"github.com/yilingfeng/apiserver/pkg/errno"
	"github.com/yilingfeng/apiserver/pkg/auth"
	"github.com/yilingfeng/apiserver/pkg/token"
)

// Login generates the authentication token
// if the password was matched with the specified account.
func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	
	// Get the user information by the login username.
	db, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	
	// Compare the login password with the user password.
	if err := auth.Compare(db.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	
	t, err := token.Sign(c, token.Context{ID: db.ID, Username: db.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	
	SendResponse(c, nil, model.Token{Token: t})
	
}
