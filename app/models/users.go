package models

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/aminshokripwa/Golang-Restful-API-using-GORM/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	//Model of table
	Id       string
	Name     string
	Username string `gorm:"type:varchar(100);unique_index"`
	Password string
	Tokens   string
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if user.Name == "" {
		return u.Message(false, "Name should be on the payload"), false
	}

	if user.Username == "" {
		return u.Message(false, "Email should be on the payload"), false
	}

	if !strings.Contains(user.Username, "@") {
		return u.Message(false, "Email address is required"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	if err != nil {
		//u.Message(false, "There was an internal error")
		return nil
	}
	user.Password = string(hash)

	//check if username exists
	result := GetUsername(user.Username)
	//fmt.Println(result)
	if result != nil {
		//fmt.Println("This username has already been used")
		return u.Message(false, "This username has already been used")
	}

	GetDB().Create(user)

	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}

func GetUser(id int) *User {
	user := &User{}
	err := GetDB().First(&user, id).Error
	if err != nil {
		return nil
	}
	return user
}

func GetUsers() []*User {
	users := make([]*User, 0)
	err := GetDB().Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return users
}

func UpdateUser(user *User, id int) (err error) {

	//if user name used by others not update
	//fmt.Println(id)
	result := GetUsername(user.Username)
	resultit := GetId(id)
	//fmt.Println(result.ID)
	//fmt.Println(result.Tokens)
	if id > 0 && result != nil {
		if id != int(result.ID) {
			fmt.Println("This username has already been used")
			return fmt.Errorf("This username has already been used")
		}
	}

	if len([]rune(user.Password)) < 1 {
		fmt.Println("new password must be entered")
		return fmt.Errorf("new password must be entered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	//new data preppers
	user.Password = string(hash)
	//safe id
	user.Id = strconv.Itoa(id)
	//safe token
	user.Tokens = resultit.Tokens

	err = GetDB().Save(user).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func DeleteUser(user *User) (err error) {
	err = GetDB().Delete(user).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

func GetUserForUpdateOrDelete(id int, user *User) (err error) {
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsername(username string) *User {
	user := &User{}
	if err := GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return user
}

func UpdateToken(username string, tokenString string) *User {
	//https://v1.gorm.io/docs/update.html
	user := &User{}
	if err := db.Model(&user).Where("username = ?", username).Update("tokens", tokenString).Error; err != nil {
		return nil
	}
	return user
}

func GetToken(tokenString string) *User {
	user := &User{}
	if err := GetDB().Where("tokens = ?", tokenString).First(&user).Error; err != nil {
		return nil
	}
	return user
}

func GetId(id int) *User {
	user := &User{}
	if err := GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		return nil
	}
	return user
}
