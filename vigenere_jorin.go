package vigenere_jorin

import (
	"strings"
)

// Alphabet contains set of possible runes: [A-Z ]
const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "

// AlphabetLen is length of Alphabet
const AlphabetLen = len(Alphabet)

// KeyPositionReset controls whether Key Position Reset is activated
const KeyPositionReset = true

// Sanitize converts lower to upper case and removes non-runes but keeps space
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

// RotateRight rotates one rune for encryption
func RotateRight(a, b rune) rune {
	ai := strings.IndexRune(Alphabet, a)
	bi := strings.IndexRune(Alphabet, b)

	return rune(Alphabet[(ai+bi)%AlphabetLen])
}

// RotateLeft rotates one rune for decryption
func RotateLeft(a, b rune) rune {
	ai := strings.IndexRune(Alphabet, a)
	bi := strings.IndexRune(Alphabet, b)

	return rune(Alphabet[(ai-bi+AlphabetLen)%AlphabetLen])
}

// RoundRight rotates one message for encryption
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

			if KeyPositionReset {
				alt = !alt
			}
		} else {
			ki = (ki + 1) % len(key)
		}
	}

	return out
}

// RoundLeft rotates one message for decryption
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

			if KeyPositionReset {
				alt = !alt
			}
		} else {
			ki = (ki + 1) % len(key)
		}
	}

	return out
}

// Encrypt encrypts a message
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

// Decrypt decrypts a message
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
