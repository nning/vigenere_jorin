package main

import (
	"fmt"
	"os"
	"strings"
)

// Set of possible runes: [A-Z ]
const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
const ALPHABET_LEN = len(ALPHABET)

// Converts lower to upper case and removes non-runes but keeps space
func Sanitize(in string) string {
	out := []rune{}

	for _, v := range in {
		if 65 <= v && v <= 90 || v == 32 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}

	return string(out)
}

func RotateRight(a, b rune) rune {
	ai := strings.IndexRune(ALPHABET, a)
	bi := strings.IndexRune(ALPHABET, b)

	return rune(ALPHABET[(ai+bi)%ALPHABET_LEN])
}

func RotateLeft(a, b rune) rune {
	ai := strings.IndexRune(ALPHABET, a)
	bi := strings.IndexRune(ALPHABET, b)

	return rune(ALPHABET[(ai-bi+ALPHABET_LEN)%ALPHABET_LEN])
}

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "%s encrypt|decrypt <key> <message>\n", os.Args[0])
		os.Exit(1)
	}

	key := Sanitize(os.Args[2])
	msg := Sanitize(os.Args[3])

	out := make([]rune, 0, len(msg))

	ki := 0
	alt := false

	if os.Args[1] == "encrypt" {
		for _, v := range msg {
			k := rune(key[ki])

			c := RotateRight(v, k)
			out = append(out, c)

			if v == ' ' && alt {
				ki = 0
				alt = !alt
			} else {
				ki = (ki + 1) % len(key)
			}
		}
	}

	if os.Args[1] == "decrypt" {
		for _, v := range msg {
			k := rune(key[ki])

			c := RotateLeft(v, k)
			out = append(out, c)

			if c == ' ' && alt {
				ki = 0
				alt = !alt
			} else {
				ki = (ki + 1) % len(key)
			}
		}
	}

	fmt.Println(string(out))
}
