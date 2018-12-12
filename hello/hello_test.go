package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Stokes", "")
		want := "Hello, Stokes"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Ellie", "Spanish")
		want := "Hola, Ellie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Florian", "French")
		want := "Bonjour, Florian"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Bishop", "Russian")
		want := "Приве́т, Bishop"
		assertCorrectMessage(t, got, want)
	})
}
