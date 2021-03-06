// package tests comment
package main

import "testing"

func runTests(tests map[int]string, t *testing.T) {
	for input, expected := range tests {
		result, err := SpellNumber(input)
		if result != expected || err != nil {
			t.Errorf("For %d, got %#q but %#q expected", input, result, expected)
		}
	}
}

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

func TestSpellNumberPrimitives(t *testing.T) {
	tests := map[int]string {
		1: "один",
		2: "два",
		3: "три",
		4: "четыре",
		5: "пять",
		6: "шесть",
		7: "семь",
		8: "восемь",
		9: "девять",

		10: "десять",
		11: "одиннадцать",
		12: "двенадцать",
		13: "тринадцать",
		14: "четырнадцать",
		15: "пятнадцать",
		16: "шестнадцать",
		17: "семнадцать",
		18: "восемнадцать",
		19: "девятнадцать",

		20: "двадцать",
		30: "тридцать",
		40: "сорок",
		50: "пятьдесят",
		60: "шестьдесят",
		70: "семьдесят",
		80: "восемьдесят",
		90: "девяносто",

		100: "сто",
		200: "двести",
		300: "триста",
		400: "четыреста",
		500: "пятьсот",
		600: "шестьсот",
		700: "семьсот",
		800: "восемьсот",
		900: "девятьсот",
	}

	runTests(tests, t)
}

func TestSpellNumber21To100(t *testing.T) {
	tests := map[int]string {
		21: "двадцать один",
		43: "сорок три",
		99: "девяносто девять",
	}

	runTests(tests, t)
}

func TestSpellNumber100To999(t *testing.T) {
	tests := map[int]string {
		194: "сто девяносто четыре",
		211: "двести одиннадцать",
		285: "двести восемьдесят пять",
		294: "двести девяносто четыре",
		569: "пятьсот шестьдесят девять",
		643: "шестьсот сорок три",
		827: "восемьсот двадцать семь",
	}

	runTests(tests, t)
}

func TestSpellNumberZero(t *testing.T) {
	tests := map[int]string { 0: "ноль" }
	runTests(tests, t)
}

func TestSpellNumberMillion(t *testing.T) {
	tests := map[int]string { 1000000: "один миллион" }
	runTests(tests, t)
}

func TestSpellNumberThousands(t *testing.T) {
	tests := map[int]string {
		1000: "одна тысяча",
		1011: "одна тысяча одиннадцать",
		1213: "одна тысяча двести тринадцать",
		1984: "одна тысяча девятьсот восемьдесят четыре",
		2000: "две тысячи",
		2012: "две тысячи двенадцать",
		5869: "пять тысяч восемьсот шестьдесят девять",
		204757: "двести четыре тысячи семьсот пятьдесят семь",
	}

	runTests(tests, t)
}
