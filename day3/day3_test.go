package day3

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution_Part1(t *testing.T) {
	s := Solution{}

	assert.Equal(t, 4, s.Part1(bytes.NewBufferString("^>v<")))
}
