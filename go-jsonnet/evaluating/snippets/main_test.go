package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainFunction(t *testing.T) {
	assert.Equal(t, "great", f("utils", "foo", `"great", "bear"`))
	assert.Equal(t, float64(1), f("utils", "foo", `1, 2`))
	assert.Equal(t, true, f("utils", "foo", `true, false`))
}
