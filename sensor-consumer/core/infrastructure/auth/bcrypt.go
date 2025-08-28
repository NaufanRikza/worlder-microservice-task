package auth

import "golang.org/x/crypto/bcrypt"

type passwordHasher struct{}

type PasswordHasher interface {
	CheckPasswordHash(hash, password string) error
}

func NewPasswordHasher() PasswordHasher {
	return &passwordHasher{}
}

func (b *passwordHasher) CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
