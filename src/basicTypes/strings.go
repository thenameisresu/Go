package basicTypes

import "fmt"

// Strings in Go are UTF-8 Encoded.
func printChars(s string) {
	fmt.Printf("Chars: ")
	for _, i := range s {
		fmt.Printf("%c ", i)
	}
	fmt.Printf("\n")
}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for _, i := range s {
		fmt.Printf("%x ", i)
	}
	fmt.Printf("\n")
}

func printChar(s string) {
	fmt.Printf("Chars: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Printf("\n")
}

func printRune(s string) {
	fmt.Printf("Runes: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func StringType() {
	str1 := "Vanshika"
	str2 := "Señor"

	fmt.Printf("String: %s\n", str1)
	printChars(str1)
	printBytes(str1)
	fmt.Printf("\n")
	printChars(str2) //this prints "Señor"
	printBytes(str2)
	printChar(str2) //this prints "SeÃ±or"
	// the Unicode code point of ñ is U+00F1 and its UTF-8 encoding occupies 2 bytes c3 and b1.
	// We are trying to print characters assuming that each code point will be one byte long which is wrong.
	// In UTF-8 encoding a code point can occupy more than 1 byte.

	// We solve this with rune
	printRune(str2)
}
