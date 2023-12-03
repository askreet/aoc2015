package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRibbonNeed(t *testing.T) {
	assert.Equal(t, 34, RibbonNeed(2, 3, 4))
	assert.Equal(t, 14, RibbonNeed(1, 1, 10))
}
