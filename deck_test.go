package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Error deck lengh is not 4, got %v", len(d))

	}
}
