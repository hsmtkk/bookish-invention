package collatz_test

import (
	"testing"

	"github.com/hsmtkk/bookish-invention/collatz"
	"github.com/stretchr/testify/assert"
)

func TestCollatzSteps(t *testing.T) {
	want := 9
	got := collatz.CollatzSteps(12)
	assert.Equal(t, want, got)
}

func TestCollatzNumbers(t *testing.T) {
	want := []int{6, 3, 10, 5, 16, 8, 4, 2, 1}
	got := collatz.CollatzNumbers(12)
	assert.Equal(t, want, got)
}
