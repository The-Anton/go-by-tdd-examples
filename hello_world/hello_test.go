package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {
		got := Hello("Charlee", "english")
		want := "Hello!, Charlee"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func (t *testing.T) {
		got := Hello("", "english")
		want := "Hello!, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func (t *testing.T) {
		got := Hello("Charlee", "spanish")
		want := "Hola!, Charlee"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}