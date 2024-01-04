package passwords

import (
	"math/rand"
)

const temporaryPasswordLength = 12

func PasswordGenerator() string {

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	length := temporaryPasswordLength
	buf := make([]rune, length)
	for i := range buf {
		buf[i] = chars[rand.Intn(len(chars))]
	}
	str := string(buf)

	return str

}
