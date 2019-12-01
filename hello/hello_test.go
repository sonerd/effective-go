package main

import "testing"

func Test_hello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", spanish)
		want := "Hola, Chris"

		assertString(got, want, t)
	})

	t.Run("say hello with default", func(t *testing.T) {
		got := hello("", eng)
		want := "Hello, World"

		assertString(got, want, t)
	})
}

func assertString(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
