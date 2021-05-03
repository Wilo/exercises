package main

import (
	"testing"

	"gotest.tools/assert"
)

var (
	palindrome_name    string = "Anna"
	no_palindrome_name string = "Beto"
	empty_text         string = ""
)

func TesIsPalindrome(t *testing.T) {
	t.Run("Is Palindrome", func(t *testing.T) {
		t.Run("When the text is palindrome", func(t *testing.T) {
			result := IsPalindrome(palindrome_name)
			assert.Equal(t, result, true)
		})

		t.Run("When the text is not palindrome", func(t *testing.T) {
			result := IsPalindrome(no_palindrome_name)
			assert.Equal(t, result, false)
		})
	})

	t.Run("Reverse Text", func(t *testing.T) {
		t.Run("When the text is correct", func(t *testing.T) {
			result := Reverse(palindrome_name)
			assert.Equal(t, result, "annA")
		})
		t.Run("When the input is empty", func(t *testing.T) {
			result := Reverse(empty_text)
			assert.Equal(t, result, "")
		})
	})
}
