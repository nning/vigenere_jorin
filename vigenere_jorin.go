package vigenere_jorin

import (
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

func RoundRight(key, msg []rune) []rune {
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

func RoundLeft(key, msg []rune) []rune {
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

func Encrypt(key, msg []rune, rounds ...int) []rune {
	rds := 1

	if len(rounds) > 0 {
		rds = rounds[0]
	}

	out := make([]rune, len(msg))
	copy(out, msg)

	for i := 0; i < rds; i++ {
		out = RoundRight(key, out)
	}

	return out
}

func Decrypt(key, msg []rune, rounds ...int) []rune {
	rds := 1

	if len(rounds) > 0 {
		rds = rounds[0]
	}

	out := make([]rune, len(msg))
	copy(out, msg)

	for i := 0; i < rds; i++ {
		out = RoundLeft(key, out)
	}

	return out
}
