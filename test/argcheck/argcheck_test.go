package argcheck

import (
	"testing"

	"github.com/mucahitkurtlar/rstrings/internal/argcheck"
)

func TestCheckHelp(t *testing.T) {
	got := argcheck.CheckHelp([]string{"-h"})
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestCheckVersion(t *testing.T) {
	got := argcheck.CheckVersion([]string{"-v"})
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
