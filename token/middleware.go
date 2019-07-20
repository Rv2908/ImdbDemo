package token

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// JWTMiddleware - middleware to handle jwt verifications
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get authorization token from header
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Write([]byte(errors.New("invalid or expired jwt token").Error()))
			return
		}
		// split jwt token and verify it
		a := strings.Split(auth, " ")
		if len(a) != 2 {
			w.Write([]byte(errors.New("invalid or expired jwt token").Error()))
			return
		}
		// call jwt.Verify method to verify jwt token
		v, err := Verify(a[1], GetAccessTokenSecret())
		if err != nil {
			fmt.Println("Error in jwt middleware", err)
			w.Write([]byte(errors.New("invalid or expired jwt token").Error()))
			return
		}
		if !v {
			w.Write([]byte(errors.New("invalid or expired jwt token").Error()))
			return
		}
		// call next handler
		next(w, r)
	})
}

// AdminMiddleware - checks if current api is only accessible by admin
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get authorization token from header
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Write([]byte(errors.New("Only admin can access this api").Error()))
			return
		}
		// split jwt token
		a := strings.Split(auth, " ")
		if len(a) != 2 {
			w.Write([]byte(errors.New("Only admin can access this api").Error()))
			return
		}
		// get payload information from token
		_, isAdmin, err := ParseToken(a[1])
		if err != nil {
			fmt.Println("Error in admin middleware:", err)
			w.Write([]byte(errors.New("Only admin can access this api").Error()))
			return
		}
		if !isAdmin {
			w.Write([]byte(errors.New("Only admin can access this api").Error()))
			return
		}
		// call next handler
		next(w, r)
	})
}
