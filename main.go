package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
)

type config struct {
	Length           int
	IncludeLowercase bool
	IncludeUppercase bool
	IncludeNumbers   bool
	IncludeSymbols   bool
}

func main() {
	var c config
	flag.BoolVar(&c.IncludeLowercase, "lower", true, "Include Lowercase Letters (a-z)")
	flag.BoolVar(&c.IncludeUppercase, "upper", true, "Include Uppercase Letters (A-Z)")
	flag.BoolVar(&c.IncludeNumbers, "number", true, "Include Numbers (0-9)")
	flag.BoolVar(&c.IncludeSymbols, "sybmols", true, "Include Symbols (!@#$%^&*()[]{}|;:/?.>,<`~)")
	flag.IntVar(&c.Length, "length", 10, "Password length")
	flag.Parse()

	var characters []byte
	// Set up array of included characters
	if c.IncludeLowercase {
		characters = append(characters, []byte("abcdefghijklmnopqrstuvwxyz")...)
	}
	if c.IncludeUppercase {
		characters = append(characters, []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")...)
	}
	if c.IncludeNumbers {
		characters = append(characters, []byte("1234567890")...)
	}
	if c.IncludeSymbols {
		characters = append(characters, []byte("!@#$%^&*()[]{}|;:/?.>,<`~)")...)
	}

	var result []byte
	for x := 0; x < c.Length; x++ {
		digit, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, characters[digit.Int64()])
	}

	fmt.Printf("%s\n", result)
}
