package user

import model "Imdb/model/user"

type User interface {
	Add(name, email, phone, password string) (*model.User, error)
}
