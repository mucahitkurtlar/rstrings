package files

import (
	"testing"

	"github.com/mucahitkurtlar/rstrings/pkg/files"
)

func TestGetfiles(t *testing.T) {
	got, _ := files.GetFiles("./")
	want := []string{"files_test.go"}
	if got[0] != want[0] {
		t.Errorf("got %q want %q", got, want)
	}
}
