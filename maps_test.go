package maps

import (
	//"testify/assert"
	"testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this just a test value"}
	t.Run("known word", func(t *testing.T) {
		got, err := dict.Search("test")
		assert.Error(t, err)
		want := "this just a test value"

		assertStrings(t, got, want)
		//assert.Equalf(t, want, got, "got %q want %q", got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("test")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Errorf("expected to get an error.")
		}
		//assert.Error(t, err)
		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
