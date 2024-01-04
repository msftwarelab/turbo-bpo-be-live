package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	tokenTest, _ := New("test@test.com")
	isvalid, _ := Decode(*tokenTest)
	assert.Equal(t, isvalid, true)

}
