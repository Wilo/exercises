package main

import (
	"fmt"
	"strings"
)

func Reverse(text string) (reversed string) {
	for _, value := range text {
		reversed = string(value) + reversed
	}
	return reversed
}

func IsPalindrome(text string) bool {
	return strings.EqualFold(strings.ToLower(text), strings.ToLower(Reverse(text)))
}

func main() {
	var text string
	fmt.Print("Give me the text to analyze: ")
	fmt.Scan(&text)
	fmt.Printf("%s Is Palindrome?: %t\n", text, IsPalindrome(text))
}
