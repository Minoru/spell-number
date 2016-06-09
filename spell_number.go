// This program takes a number in a range from 0 to a million and spells it out
// in Russian.
package main

import "fmt"

// SpellNumber spells out given int in Russian.
func SpellNumber(i int) (result string, err error) {
	switch {
	// error cases
	case i < 0: err = fmt.Errorf("%d: number is too small!", i)
	case i > 1000000: err = fmt.Errorf("%d: number is too large!", i)

	//corner cases
	case i == 0: result = "ноль"
	case i == 1000000: result = "один миллион"

	default: result, err = helper(i, false)
	}
	return
}

// Concatenates spelling of a number to result only if the spelling is
// non-empty (i.e. if number is non-zero). Propagates existing errors, i.e.
// will return immediately if err != nil.
func addComponent(number int, result string, err error, isThousands bool) (output string, out_err error) {
	if err != nil {
		out_err = err
	} else {
		str, spellErr := helper(number, isThousands)
		if spellErr != nil {
			out_err = spellErr
		} else {
			output = result
			if str != "" {
				if output != "" {
					output += " "
				}
				output += str
			}

			if isThousands {
				// Just like primitives, "thousands" word
				// changes with the last digit. NB: since
				// isThousands is true, number is amount of
				// thousands, so there's no need to divide by
				// 1000 or anything.
				digit := number % 10
				if digit == 1 {
					output += " тысяча"
				} else if digit > 1 && digit < 5 {
					output += " тысячи"
				} else {
					output += " тысяч"
				}
			}
		}
	}

	return
}

// Primitives are member of the set of Russian words that are used when
// spelling words. When spelling thousands, some primitives change, e.g.
// "один"/"одна" in "двадцать-один" vs. "одна тысяча".

type Primitive struct {
	ordinary, thousands string
}

var primitives = map[int]Primitive {
	0: Primitive{},

	1: Primitive{ordinary: "один", thousands: "одна"},
	2: Primitive{ordinary: "два", thousands: "две"},
	3: Primitive{ordinary: "три"},
	4: Primitive{ordinary: "четыре"},
	5: Primitive{ordinary: "пять"},
	6: Primitive{ordinary: "шесть"},
	7: Primitive{ordinary: "семь"},
	8: Primitive{ordinary: "восемь"},
	9: Primitive{ordinary: "девять"},

	10: Primitive{ordinary: "десять"},
	11: Primitive{ordinary: "одиннадцать"},
	12: Primitive{ordinary: "двенадцать"},
	13: Primitive{ordinary: "тринадцать"},
	14: Primitive{ordinary: "четырнадцать"},
	15: Primitive{ordinary: "пятнадцать"},
	16: Primitive{ordinary: "шестнадцать"},
	17: Primitive{ordinary: "семнадцать"},
	18: Primitive{ordinary: "восемнадцать"},
	19: Primitive{ordinary: "девятнадцать"},

	20: Primitive{ordinary: "двадцать"},
	30: Primitive{ordinary: "тридцать"},
	40: Primitive{ordinary: "сорок"},
	50: Primitive{ordinary: "пятьдесят"},
	60: Primitive{ordinary: "шестьдесят"},
	70: Primitive{ordinary: "семьдесят"},
	80: Primitive{ordinary: "восемьдесят"},
	90: Primitive{ordinary: "девяносто"},

	100: Primitive{ordinary: "сто"},
	200: Primitive{ordinary: "двести"},
	300: Primitive{ordinary: "триста"},
	400: Primitive{ordinary: "четыреста"},
	500: Primitive{ordinary: "пятьсот"},
	600: Primitive{ordinary: "шестьсот"},
	700: Primitive{ordinary: "семьсот"},
	800: Primitive{ordinary: "восемьсот"},
	900: Primitive{ordinary: "девятьсот"},
}

// Spells any number in [1; 999] range. NB: zero is handled higher up the call
// stack, in SpellNumber.
func helper(i int, isThousands bool) (result string, err error) {
	if _, ok := primitives[i]; ok {
		prim := primitives[i]
		if isThousands && prim.thousands != "" {
			result = prim.thousands
		} else {
			result = prim.ordinary
		}
	} else if i > 999 {
		thousands := (i - i%1000) / 1000
		result, err = addComponent(thousands, result, err, true)
		result, err = addComponent(i%1000, result, err, false)
	} else {
		digit := i % 10
		dozens := i % 100 - digit
		hundreds := i % 1000 - dozens - digit

		result, err = addComponent(hundreds, result, err, false)
		if _, ok := primitives[dozens + digit]; ok {
			result, err = addComponent(dozens+digit, result, err, false)
		} else {
			result, err = addComponent(dozens, result, err, false)
			result, err = addComponent(digit, result, err, false)
		}
	}

	return
}

func main() {
	var input int
	fmt.Scanf("%d", &input)

	result, err := SpellNumber(input)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
