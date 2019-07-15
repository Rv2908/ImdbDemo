package user

import (
	model "Imdb/model/user"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return User{db}
}

func (u User) Add(user *model.User) (*model.User, error) {

	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if errPassword != nil {
		return nil, errPassword
	}
	sqlStatement := `INSERT INTO users (created_at, updated_at,first_name, last_name, email, password, role_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := u.db.Exec(sqlStatement, time.Now(), time.Now(), user.FirstName, user.LastName, user.Email, string(hashedPassword), user.RoleID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
