package model

// User Details
type User struct {
	ID        uint
	FirstName string
	LastName  string
}

// UserRole Specifies users role
type UserRole struct {
	ID   uint
	Role string
}

//UserLoginDetail Store User login details
type UserLoginDetail struct {
	ID       uint
	Email    string
	Password string
}

//UserCredential stores the auth token of the user
type UserCredential struct {
	ID    uint
	Token string
}
