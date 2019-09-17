package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncorrectString(t *testing.T) {
	assert.Equal(t, "incorrect string", Unpack("45"))
}

func TestValidString(t *testing.T) {
	assert.Equal(t, "aaaabccddddde", Unpack("a4bc2d5e"))
	assert.Equal(t, "abcd", Unpack("abcd"))
	assert.Equal(t, "abcdddd", Unpack("abcd4"))
}

func TestEscapedString(t *testing.T) {
	assert.Equal(t, `qwe45`, Unpack(`qwe\4\5`))
	assert.Equal(t, `qwe44444`, Unpack(`qwe\45`))
	assert.Equal(t, `qwe\\\\\`, Unpack(`qwe\\5`))
}
