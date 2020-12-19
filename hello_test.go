package main

import "testing"

func TestHello(t *testing.T) {

	assetMessage := func(t *testing.T, got string, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nam", "")
		want := "hello Nam"
		assetMessage(t, got, want)
	})

	t.Run("saying hello world, when string is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "hello world"
		assetMessage(t, got ,want)
	})

	t.Run("in vietnamese", func(t *testing.T) {
		got := Hello("Nam", vietnamese)
		want := "xin chao Nam"
		assetMessage(t, got ,want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("David", french)
		want := "Bonjour David"
		assetMessage(t, got ,want)
	})
}
