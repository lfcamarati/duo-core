package security

import (
	"golang.org/x/crypto/bcrypt"
)

func NewDefaultPasswordEncrypt() PasswordEncrypt {
	return &BCryptPasswordEncrypt{}
}

type PasswordEncrypt interface {
	Encrypt(password string) string
	CheckEncrypt(password string, hash string) bool
}

type BCryptPasswordEncrypt struct {
}

func (p *BCryptPasswordEncrypt) Encrypt(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (p *BCryptPasswordEncrypt) CheckEncrypt(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
