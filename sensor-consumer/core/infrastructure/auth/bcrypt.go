package auth

import "golang.org/x/crypto/bcrypt"

type passwordHasher struct{}

type PasswordHasher interface {
	CheckPasswordHash(password, hash string) error
}

func NewPasswordHasher() PasswordHasher {
	return &passwordHasher{}
}

func (b *passwordHasher) CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
