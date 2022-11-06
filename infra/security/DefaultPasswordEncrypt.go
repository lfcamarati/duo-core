package security

import (
	"golang.org/x/crypto/bcrypt"
)

func NewDefaultPasswordEncrypt() PasswordEncrypt {
	return &PasswordEncryptSha256{}
}

type PasswordEncrypt interface {
	Encrypt(password string) string
	CheckEncrypt(password string, hash string) bool
}

type PasswordEncryptSha256 struct {
}

func (p *PasswordEncryptSha256) Encrypt(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (p *PasswordEncryptSha256) CheckEncrypt(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
