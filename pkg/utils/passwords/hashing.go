package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordHashAndSalt(pwd []byte) (*string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	h := string(hash)
	return &h, nil
}
