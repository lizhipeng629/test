package test_study

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hell, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
