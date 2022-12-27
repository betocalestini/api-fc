package entity

import (
	"github.com/betocalestini/api-fc/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}

func (u *User) Validateassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
