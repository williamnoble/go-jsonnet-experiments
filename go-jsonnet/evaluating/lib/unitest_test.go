package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainFunction(t *testing.T) {
	testFile := "utils"
	testFn := "foo"
	assert.Equal(t, "great", F(testFile, testFn, `"great", "bear"`))
	assert.Equal(t, float64(1), F(testFile, testFn, `1, 2`))
	assert.Equal(t, true, F(testFile, testFn, `true, false`))
}
