// package tests comment
package main

import "testing"

func TestSpellNumber(t *testing.T) {
	// Should return error if input is too large
	_, err := SpellNumber(1000001)
	if err == nil {
		t.Error("Didn't fail on 1000001")
	}
}
