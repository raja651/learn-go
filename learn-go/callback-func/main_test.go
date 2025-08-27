package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	total := add(5, 7)

	if total != 12 {
		log.Fatalf("error - want %v and got %v", 12, total)
	}
}

func TestParadise(t *testing.T) {
	got := paradise("Bali")
	want := "My idea of paradise is: Baali"
	if got != want {
		log.Fatalf("error - want %s and got %s", want, got)
	}
}
