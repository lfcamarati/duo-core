package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Token struct {
	Username string
}

func GenerateJWT(username string) (*string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	claims["sub"] = username
	claims["authorized"] = true

	var secretKey = []byte(viper.GetString("jwt.secretKey"))
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func VerifyJWT(tokenString string) error {
	if tokenString == "" {
		return errors.New("not Authorized")
	}

	_, err := parse(tokenString)

	if err != nil {
		return err
	}

	return nil
}

func DecodeJwt(tokenJwt string) (*Token, error) {
	parsedToken, err := parse(tokenJwt)

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if ok && parsedToken.Valid {
		userToken := Token{
			Username: claims["sub"].(string),
		}

		return &userToken, nil
	}

	return nil, errors.New("error decoding token")
}

func parse(tokenJwt string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenJwt, func(tokenJwt *jwt.Token) (interface{}, error) {
		if _, ok := tokenJwt.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("not Authorized")
		}

		var secretKey = []byte(viper.GetString("jwt.secretKey"))
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
