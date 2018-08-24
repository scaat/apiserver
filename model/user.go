package model

import (
	"github.com/yilingfeng/apiserver/pkg/constvar"
	"fmt"
	"github.com/yilingfeng/apiserver/pkg/auth"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

// Create creates a new user account
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.ID = id
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Self.Model(&u).Update(map[string]interface{}{"username": u.Username, "password": u.Password}).Error
	// return DB.Self.Save(u).Error
}

// GetUser gets an user by the username.
func GetUser(username string) (*UserModel, error) {
	user := UserModel{}
	db := DB.Self.Where("username = ?", username).First(&user)
	return &user, db.Error
}

// ListUser lists all users.
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	
	users := make([]*UserModel, 0)
	var count uint64
	
	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}
	
	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}
	
	return users, count, nil
}

// Compare compares with the plain text password. Returns true if it's the same as the encrypted one (in the `user` struct)
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt encrypts the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate validates the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
