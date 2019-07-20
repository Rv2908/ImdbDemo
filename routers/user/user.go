package user

import (
	interfaces "Imdb/interfaces/user"
	user "Imdb/model/user"
	token "Imdb/token"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User contains the instance of user controller and logger instance
type User struct {
	userController interfaces.User
	logger         *log.Logger
}

//Logger log all the incoming method to user route
func (u User) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer u.logger.Printf("New Request in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

//NewUserRouter Return instance of User Router
func NewUserRouter(userController interfaces.User, logger *log.Logger) User {
	return User{
		userController: userController,
		logger:         logger,
	}
}

//Register This method consist of all the routes for the user
func (u User) Register(s *http.ServeMux) {
	s.HandleFunc("/user", u.Logger(u.addUser))
	s.HandleFunc("/user/signin", u.Logger(u.signin))
}

func (u User) addUser(w http.ResponseWriter, r *http.Request) {
	user := &user.User{}
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

func (u User) signin(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type Response struct {
		User        user.UserResponse
		AccessToken string
	}

	request := &Request{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userInfo, err := u.userController.Signin(request.Email)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No Such User Exist"))
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(request.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessToken, err := token.GenerateAccessToken(request.Email, userInfo.Role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response Response

	response.User = userInfo
	response.AccessToken = accessToken

	b, _ := json.Marshal(response)
	w.Write(b)
}
