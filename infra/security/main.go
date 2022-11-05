package security

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secret")

func GenerateJWT(username string) (*string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	claims["authorized"] = true
	claims["user"] = username

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func VerifyJWT(tokenString string) error {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	if tokenString == "" {
		return errors.New("not Authorized")
	}

	_, err := jwt.Parse(tokenString, func(tokenJwt *jwt.Token) (interface{}, error) {
		if _, ok := tokenJwt.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("not Authorized")
		}

		return secretKey, nil
	})

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
