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
	prefix := input[0:2]
	check := input[2:4]
	bban := input[4:]
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
	mod := iban.Mod(&iban, big.NewInt(97)).Int64()
	if mod != 1 {
		if check == "00" {
			fmt.Printf("Generated IBAN with checksum replaced: %s%02d%s\n", prefix, 98-mod, bban)
		} else {
			fmt.Println("IBAN incorrect or input error: ", 98-mod)
		}
		os.Exit(int(98-mod))
	}
	fmt.Println("IBAN probably correct")
}
