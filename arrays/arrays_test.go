package arrays

import (
	"testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Using testify/assert :", func(t *testing.T) {
		assert := assert.New(t)
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assert.Equalf(want, got, "got %d want %d given, %v", got, want, numbers)
	})

	t.Run("Using standard testing :", func(t *testing.T) {

	})
}

func TestSumAssert(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
