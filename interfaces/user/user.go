package user

import model "Imdb/model/user"

type User interface {
	Add(user *model.User) (*model.User, error)
}
