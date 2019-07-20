package token

import "time"

// AccessTokenExpiryTime - access token expiry time in seconds
const AccessTokenExpiryTime = 1800 // 60 * 30 minutes 1800 seconds
var accessTokenSecret string

// SetAccessTokenSecret - generates and sets a random secret token for generating access keys
func SetAccessTokenSecret() {
	accessTokenSecret = GeneratePrivateKey()
}

// GetAccessTokenSecret - returns the access token key
func GetAccessTokenSecret() string {
	return accessTokenSecret
}

// GenerateAccessToken - generates and returns new access token
func GenerateAccessToken(email string, IsAdmin bool) (string, error) {
	// create jwt header
	// for our use case, we'll be using SHA256 algorithm and type JWT
	header := Header{
		Typ: "JWT",
		Alg: "SHA256",
	}
	// create payload information
	payload := Payload{
		Email:     email,
		IsAdmin:   IsAdmin,
		ExpiredAt: time.Now().Unix() + AccessTokenExpiryTime,
	}
	// generate and return new access token
	return GetJwtToken(header, payload, accessTokenSecret)
}
