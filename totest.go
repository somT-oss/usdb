package main

import "strconv"

func Fooer(input int) string {
	isFoo := input % 3 == 0

	if isFoo {
		return "Foo"
	}
	return strconv.Itoa(input)
} 