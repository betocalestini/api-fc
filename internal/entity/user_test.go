package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUser(t *testing.T) {
	user, err := NewUser("Roberto Calestini", "betocalestini@hotmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.ID)
	assert.NotNil(t, user.Password)
	assert.Equal(t, "Roberto Calestini", user.Name)
	assert.Equal(t, "betocalestini@hotmail.com", user.Email)
}

func Fuzz_NewUser(f *testing.F) {
	// Add test cases with specific values for name, email, and password
	f.Add("Alice", "alice@example.com", "123456789")

	f.Fuzz(func(t *testing.T, name, email, password string) {
		user, err := NewUser(name, email, password)

		if name == "" || email == "" || password == "" {
			assert.NotNil(t, err)
			return
		}
		if assert.Nil(t, err) {
			assert.NotNil(t, user)
			assert.NotNil(t, user.ID)
			assert.NotNil(t, user.Password)
			assert.Equal(t, name, user.Name)
			assert.Equal(t, email, user.Email)
		}

	})
}

func Test_ValidatePassword(t *testing.T) {
	user, err := NewUser("Roberto Calestini", "betocalestini@hotmail.com", "123asd456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123asd456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "1234656", user.Password)
}
