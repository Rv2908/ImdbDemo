package token

import (
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	b64 "encoding/base64"

	"crypto/hmac"
	"crypto/sha256"
)

// Payload - provides structure of information contained inside the jwt token
type Payload struct {
	Email     string
	ExpiredAt int64 // time in seconds
	IsAdmin   bool
}

// Header - provides header for the jwt token
type Header struct {
	Alg string
	Typ string
}

// GeneratePrivateKey - generates new private key
func GeneratePrivateKey() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var b []byte
	for i := 0; i < 5; i++ {
		b = append(b, byte(r.Intn(122-65)+65))
	}
	return hex.EncodeToString(b)
}

// GetJwtToken - creates a jwt token
func GetJwtToken(header Header, payload Payload, secret string) (string, error) {
	// create header for jwt token
	h, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	// base64 encode the header
	b64Header := b64.StdEncoding.EncodeToString([]byte(h))
	// create payload for jwt
	p, err := json.Marshal(payload)
	// base64 encode the payload
	b64Payload := b64.StdEncoding.EncodeToString([]byte(p))
	if err != nil {
		return "", err
	}
	// encode HMAC SHA256 header and payload using private key to generate JWT token
	hm := hmac.New(sha256.New, []byte(secret))
	hm.Write([]byte(b64Header + "." + b64Payload + secret))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(hm.Sum(nil))
	return b64Header + "." + b64Payload + "." + sha, nil
}

// Verify - verifies the given jwt token and returns email from payload
func Verify(token, secret string) (bool, error) {
	if token == "" {
		return false, nil
	}
	// split token to get header, payload and signature
	a := strings.Split(token, ".")
	if len(a) != 3 {
		return false, nil
	}
	payloadString := a[1]
	// decode payload
	p, err := b64.StdEncoding.DecodeString(payloadString)
	if err != nil {
		return false, err
	}
	var payload Payload
	if err := json.Unmarshal(p, &payload); err != nil {
		return false, err
	}
	// check if token is expired
	if payload.ExpiredAt < time.Now().Unix() {
		return false, nil
	}
	// generate and verify hmac
	hm := hmac.New(sha256.New, []byte(secret))
	hm.Write([]byte(a[0] + "." + a[1] + secret))
	expected := hm.Sum(nil)
	actual, err := hex.DecodeString(a[2])
	if err != nil {
		return false, nil
	}
	verified := hmac.Equal(actual, expected)
	return verified, nil
}

// ParseToken - parses the given token and returns email and isAdmin flag from token
func ParseToken(token string) (string, bool, error) {
	if token == "" {
		return "", false, nil
	}
	// split token to get header, payload and signature
	a := strings.Split(token, ".")
	if len(a) != 3 {
		return "", false, nil
	}
	payloadString := a[1]
	// decode payload
	p, err := b64.StdEncoding.DecodeString(payloadString)
	if err != nil {
		return "", false, err
	}
	var payload Payload
	if err := json.Unmarshal(p, &payload); err != nil {
		return "", false, err
	}
	return payload.Email, payload.IsAdmin, nil
}
