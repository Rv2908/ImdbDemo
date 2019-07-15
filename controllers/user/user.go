package user

import(
	"fmt"
	model "Imdb/model/user"
) 

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Add(name, email, phone, password string) (*model.User,error) {
	fmt.Println("Hello World!")
	return nil,nil
}
