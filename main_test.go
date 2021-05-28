package main

import ( 
	"testing"
)


func TestMain(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func (t *testing.T) {
		got := Hello("Tau", "spanish")
		want := "Hola, Tau"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello, world when empty string is supplied", func(t *testing.T){
		got := Hello("", "spanish")
		want := "Hola, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
}
