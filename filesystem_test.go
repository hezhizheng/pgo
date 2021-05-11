package pgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileExists(t *testing.T) {
	path := "./filesystem_test.go"
	exists := FileExists(path)
	assert.Equal(t, true, exists)
}