package user

import (
	user "Imdb/model/user"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User it contains the instance of database
type User struct {
	db *sql.DB
}

//NewUser this will return the instance of user with database instance in it
func NewUser(db *sql.DB) User {
	return User{db}
}

//Add this will create a new user
func (u User) Add(user *user.User) (*user.User, error) {

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

//Signin Login the user and send the authentication token
func (u User) Signin(EmailID string) (user.UserResponse, error) {
	type UserInfo struct {
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	var userInfo user.UserResponse

	// Get the existing entry present in the database for the given username
	sqlStatement := `select 
					ur.first_name,
					ur.last_name,
					ur.email,
					ur.password ,
					case 
						when rl.role = 'Admin' then true
						else false
					end as role
					from users ur
					left join roles rl on ur.role_id = rl.id
					where ur.email like $1`

	if err := u.db.QueryRow(sqlStatement, EmailID).Scan(&userInfo.FirstName,
		&userInfo.LastName,
		&userInfo.Email,
		&userInfo.Password,
		&userInfo.Role); err != nil {
		return userInfo, err
	}

	return userInfo, nil
}
