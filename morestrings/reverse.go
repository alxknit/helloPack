// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

import "fmt"

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	fmt.Println("Other module")
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseRunes2(s string) string {
	fmt.Println("New func")
	r := []rune(s)

	z := []rune(s)
	counter := 0
	for i := len(r) - 1; i > 0; i-- {
		z[counter] = r[i]
		counter++

	}
	return string(z)
}
