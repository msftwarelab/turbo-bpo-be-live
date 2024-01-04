package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimObjectChar(t *testing.T) {
	trimedTest := TrimObjectChar("/test/")
	assert.Equal(t, trimedTest, "test")
	trimedTest = TrimObjectChar("**/test/")
	assert.Equal(t, trimedTest, "**test")

}
