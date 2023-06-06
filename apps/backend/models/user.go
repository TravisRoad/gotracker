package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Role     string
}

func (u *User) Save() (*User, error) {
	err := DB.Save(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func (u *User) PrepareTO() {
	u.Password = ""
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(uid uint) (*User, error) {
	var user User
	if err := DB.First(&user, uid).Error; err != nil {
		return &User{}, err
	}

	return &user, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string) (uint, error) {
	var err error
	u := User{}
	err = DB.Where("username = ?", username).First(&u).Error
	if err != nil {
		return 0, err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil {
		return 0, err
	}

	return u.ID, nil
}
