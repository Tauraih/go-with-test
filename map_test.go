package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "another test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "another test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

// func TestAdd(t *testing.T) {
// 	dicttionary := Dictionary{}
// 	dicttionary.Add("test", "just a test")

// 	want := "just a test"
// 	got, err := dicttionary.Search("test")
// 	if err != nil {
// 		t.Fatal("should find added word:", err)
// 	}

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
