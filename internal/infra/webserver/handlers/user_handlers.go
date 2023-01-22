package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/betocalestini/api-fc/internal/dto"
	"github.com/betocalestini/api-fc/internal/entity"
	"github.com/betocalestini/api-fc/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.UserDB.Create(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
