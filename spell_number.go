// This program takes a number in a range from 0 to a million and spells it out
// in Russian.
package main

import (
	"fmt"
	"errors"
)

// SpellNumber spells out given int in Russian.
func SpellNumber(i int) (result string, err error) {
	primitives := map[int]string {
		0: "",

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

	if i < 0 {
		err = errors.New("Number is too small!")
	} else if i > 1000000 {
		err = errors.New("Number is too large!")
	} else if _, ok := primitives[i]; ok {
		result = primitives[i]
	} else {
		// Concatenates spelling of a number to result only if the
		// spelling is non-empty (i.e. if number is non-zero).
		// Propagates existing errors, i.e. will return immediately if
		// err != nil.
		addComponent := func(number int, result string, err error) (output string, out_err error) {
			if err != nil {
				out_err = err
			} else {
				str, spellErr := SpellNumber(number)
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
				}
			}
			return
		}

		digit := i % 10
		dozens := i % 100 - digit
		hundreds := i % 1000 - dozens - digit

		result, err = addComponent(hundreds, result, err)
		result, err = addComponent(dozens, result, err)
		result, err = addComponent(digit, result, err)

		if err != nil {
			err = fmt.Errorf("No implementaton for input %d", i)
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
