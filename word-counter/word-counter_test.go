package wordcounter

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("count words if lines flag is false", func(t *testing.T) {
		input := bytes.NewBufferString("word1 word2 word3 word4\n")
		want := 4
		got := Count(input, false, false)
		assert(t, got, want)

	})
	t.Run("count lines if lines flag is true", func(t *testing.T) {
		input := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
		want := 3
		got := Count(input, true, false)
		assert(t, got, want)
	})

	t.Run("count bytes if bytes flag is true", func(t *testing.T) {
		input := bytes.NewBufferString("hello")
		want := 5
		got := Count(input, false, true)
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %d, got %d instead\n", want, got)
	}
}
