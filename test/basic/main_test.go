package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	assert.Equal(t, AddOne(1), 2, "abc")
}