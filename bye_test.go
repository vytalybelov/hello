package hello

import "testing"

func TestBye(t *testing.T) {
    want := "Bye"
    if got := Bye(); got != want {
        t.Errorf("Bye() = %q, want %q", got, want)
    }
}
