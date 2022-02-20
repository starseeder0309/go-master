package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var result string

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			result += string('A' + (c - 'a'))
		} else {
			result += string(c)
		}
	}

	return result
}

func ToUpper2(str string) string {
	var builder strings.Builder

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}

	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
