package rstrings

import (
	"os"
	"testing"

	"github.com/mucahitkurtlar/rstrings/internal/rstrings"
)

func TestRstrings(t *testing.T) {
	want, _ := os.ReadFile("./strings_rstrings_test_go.txt")
	got, _ := rstrings.Rstrings(".")
	if got[0] != string(want) {
		t.Errorf("got %q want %q", got, want)
	}
}
