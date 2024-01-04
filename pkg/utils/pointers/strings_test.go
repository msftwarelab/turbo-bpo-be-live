package pointers_test

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	s := "hello world"
	result := pointers.String(s)
	assert.NotNil(t, result)
	assert.Equal(t, s, *result)
}
