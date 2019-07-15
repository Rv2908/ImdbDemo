package user

import (
	interfaces "Imdb/interfaces/user"
	"net/http"
)

type User struct {
	userController interfaces.User
}

func NewUserRouter(userController interfaces.User) User {
	return User{
		userController: userController,
	}
}

func (u User) Register(s *http.ServeMux) {
	s.HandleFunc("/user", u.addUser)
}

func (u User) addUser(w http.ResponseWriter, r *http.Request) {
	u.userController.Add("", "", "", "")
}
