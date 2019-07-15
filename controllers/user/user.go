package user

import(
	"fmt"
	model "Imdb/model/user"
	"database/sql"
) 

type User struct{
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return User{db}
}

func (u User) Add(name, email, phone, password string) (*model.User,error) {
	fmt.Println("Hello World!")
	return nil,nil
}
