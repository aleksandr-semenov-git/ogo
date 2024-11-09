package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hi()
	want := "Let's go for a walk, after we drink tea."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
