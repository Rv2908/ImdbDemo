package user

import (
	interfaces "Imdb/interfaces/user"
	model "Imdb/model/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	userController interfaces.User
	logger         *log.Logger
}

func (u User) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer u.logger.Printf("New Request in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func NewUserRouter(userController interfaces.User, logger *log.Logger) User {
	return User{
		userController: userController,
		logger:         logger,
	}
}

func (u User) Register(s *http.ServeMux) {
	s.HandleFunc("/user", u.Logger(u.addUser))
}

func (u User) addUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	if _, err := u.userController.Add(user); err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("User Created Succcessfully"))

}
