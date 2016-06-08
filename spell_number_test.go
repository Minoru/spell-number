// package tests comment
package main

import "testing"

func TestSpellNumberTooLarge(t *testing.T) {
	// Should return error if input is too large
	_, err := SpellNumber(1000001)
	if err == nil {
		t.Error("Didn't fail on 1000001")
	}
}

func TestSpellNumberTooSmall(t *testing.T) {
	// Should return error if input is too small
	_, err := SpellNumber(-10)
	if err == nil {
		t.Error("Didn't fail on -10")
	}
}
