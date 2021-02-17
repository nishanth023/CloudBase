package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwt struct{}

func (j *Jwt) key() []byte {}

func (j *Jwt) Generate(userid string, role int8) (AccessToken string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = role
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	AccessToken, err = token.SignedString(j.key())
	if err != nil {
		return "", err
	}
	return AccessToken, nil
}

func (j *Jwt) Verify(tokenString string) (payload *jwt.Token, err error) {

	payload, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.key(), nil
	})

	return
}
