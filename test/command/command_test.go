package command

import (
	"testing"

	"github.com/mucahitkurtlar/rstrings/pkg/command"
)

func TestRun(t *testing.T) {
	got, _ := command.Run("go", "version")
	want := "go version"
	if got[0:10] != want {
		t.Errorf("got %q want %q", got, want)
	}
}
