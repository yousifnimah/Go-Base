package JWT

import (
	"base/App/Handlers/Redis"
	"base/App/Models"
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

var Claims *Claim

type Claim struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	jwt.StandardClaims
}

func GenerateJWT(User *Models.User, expirationTime time.Time) (tokenString string, err error) {
	claims := &Claim{
		ID:       User.ID,
		Username: User.Username,
		FullName: User.FullName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	err = Redis.SetAccess(claims.ID, claims.Username, expirationTime, tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ValidateJWT(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	Claims = claims
	return
}
