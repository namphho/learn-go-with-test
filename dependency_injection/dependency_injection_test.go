package main

import (
	"bytes"
	"testing"
)

func testGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nam")

	got := buffer.String()
	want := "Hello, Nam"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
