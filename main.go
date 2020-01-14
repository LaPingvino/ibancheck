package main

import (
	"fmt"
	"math/big"
	"os"
)

func let(c rune) string {
	n := -1
	switch {
	case c <= 'Z' && c >= 'A':
		n = int(c) - 'A' + 10
	case c <= 'z' && c >= 'a':
		n = int(c) - 'a' + 10
	case c >= '0' && c <= '9':
		n = int(c) - '0'
	}
	if n > -1 {
		return fmt.Sprint(n)
	}
	return ""
}

func main() {
	var input, digits, first string
	var iban big.Int
	fmt.Print("Enter an IBAN number: ")
	fmt.Scanln(&input)
	for i, c := range input {
		if i < 4 {
			first += let(c)
			continue
		}
		digits += let(c)
	}
	digits += first
	fmt.Println(digits)
	iban.SetString(digits, 10)
	if iban.Mod(&iban, big.NewInt(97)).Int64() != 1 {
		fmt.Println("IBAN incorrect or input error")
		os.Exit(97)
	}
	fmt.Println("IBAN probably correct")
}
