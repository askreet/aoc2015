package day1

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution_Part1_Examples(t *testing.T) {
	s := Solution{}

	assert.Equal(t, 0, s.Part1(bytes.NewBufferString("(())")))
	assert.Equal(t, 0, s.Part1(bytes.NewBufferString("()()")))

	assert.Equal(t, 3, s.Part1(bytes.NewBufferString("(((")))
	assert.Equal(t, 3, s.Part1(bytes.NewBufferString("(()(()(")))
	assert.Equal(t, 3, s.Part1(bytes.NewBufferString("))(((((")))

	assert.Equal(t, -1, s.Part1(bytes.NewBufferString("())")))
	assert.Equal(t, -1, s.Part1(bytes.NewBufferString("))(")))

	assert.Equal(t, -3, s.Part1(bytes.NewBufferString(")))")))
	assert.Equal(t, -3, s.Part1(bytes.NewBufferString(")())())")))
}
