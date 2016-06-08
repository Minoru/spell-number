// This program takes a number in a range from 0 to a million and spells it out
// in Russian.
package main

import (
	"fmt"
	"errors"
)

// SpellNumber spells out given int in Russian.
func SpellNumber(i int) (result string, err error) {
	if i > 1000000 {
		err = errors.New("Number is too large!")
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
