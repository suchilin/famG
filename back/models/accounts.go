package models

import (
	"github.com/jinzhu/gorm"
	u "backend/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

//a struct to rep user account
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token";sql:"-"`
	Password string `json:"password,omitempty"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(422, "Email address is required"), false
	}
	if len(account.Username) < 0 {
		return u.Message(422, "Username is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(422, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(422, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(422, "Email address already in use by another user."), false
	}

	return u.Message(200, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(422, "Failed to create account, connection error.")
	}
	// response := u.Message(200, "Account has been created")
	// response["account"] = account
	return u.Message(200, "Account created succefull")
}

func Login(username, password string) (map[string]interface{}, bool) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", username).Or("username = ?", username).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(422, "This account doesn't exist"), false
		}
		return u.Message(422, "Connection error. Please retry"), false
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(422, "Invalid login credentials. Please try again"), false
	}

	return u.Message(200, "Login succefull"), true
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
