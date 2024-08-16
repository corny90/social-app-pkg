package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/nwtgck/go-fakelish"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/big"
	"time"
)

// GenerateUniqueUsername generates a username with predefined settings and appends a Unix timestamp.
func GenerateUniqueUsername() string {
	minLength := 6          // Minimum length of each word
	maxLength := 12         // Maximum length of each word
	wordCount := 2          // Number of words to generate
	separator := ""         // Separator between words
	includeNumbers := false // Include numbers

	initialString := ""

	for i := 0; i < wordCount; i++ {
		// Generate a random length between minLength and maxLength
		randomLength, _ := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength+1)))
		length := int(randomLength.Int64()) + minLength

		// Generate a fake word with the random length
		fakeWord := fakelish.GenerateFakeWord(length, length)
		// Capitalize the first letter
		fakeWord = cases.Title(language.AmericanEnglish).String(fakeWord)

		// Append the separator if it's not the first word
		if i != 0 {
			initialString += separator
		}

		// Append the fake word
		initialString += fakeWord
	}

	// If the user wants numbers
	if includeNumbers {
		// Loop to append four random digits
		for i := 0; i < 4; i++ {
			// Random integer between 0 and 9
			randomNumber, _ := rand.Int(rand.Reader, big.NewInt(10))
			// Append the number to the initial string
			initialString += fmt.Sprintf("%d", randomNumber)
		}
	}

	// Append a Unix timestamp
	timestamp := time.Now().Unix()
	initialString += fmt.Sprintf("-%d", timestamp)

	return initialString
}
