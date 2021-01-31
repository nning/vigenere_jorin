package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Set of possible runes: [A-Z ]
const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
const ALPHABET_LEN = len(ALPHABET)

// Converts lower to upper case and removes non-runes but keeps space
func Sanitize(in string) []rune {
	out := []rune{}

	for _, v := range in {
		if 65 <= v && v <= 90 || v == 32 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}

	return out
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

func Encrypt(key []rune, msg []rune) []rune {
	out := make([]rune, 0, len(msg))

	ki := 0
	alt := false

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

	return out
}

func Decrypt(key []rune, msg []rune) []rune {
	out := make([]rune, 0, len(msg))

	ki := 0
	alt := false

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

	return out
}

func main() {
	if len(os.Args) != 4 && len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "%s encrypt|decrypt <key> <message> [rounds]\n", os.Args[0])
		os.Exit(1)
	}

	key := Sanitize(os.Args[2])
	msg := Sanitize(os.Args[3])

	rounds := 1

	if len(os.Args) == 5 {
		rounds, _ = strconv.Atoi(os.Args[4])
	}

	out := make([]rune, len(msg))
	copy(out, msg)

	if os.Args[1] == "encrypt" {
		for i := 0; i < rounds; i++ {
			out = Encrypt(key, out)
		}
	}

	if os.Args[1] == "decrypt" {
		for i := 0; i < rounds; i++ {
			out = Decrypt(key, out)
		}
	}

	fmt.Println(string(out))
}
