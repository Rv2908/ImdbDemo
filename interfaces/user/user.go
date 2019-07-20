package user

import user "Imdb/model/user"

//User it defines the behaviour of user table
type User interface {
	Add(user *user.User) (*user.User, error)
	Signin(EmailID string) (user.UserResponse, error)
}
