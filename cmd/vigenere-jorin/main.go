package main

import (
	"fmt"
	"nning.io/go/vigenere_jorin"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 && len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "%s encrypt|decrypt <key> <message> [rounds]\n", os.Args[0])
		os.Exit(1)
	}

	key := vigenere_jorin.Sanitize(os.Args[2])
	msg := vigenere_jorin.Sanitize(os.Args[3])

	rounds := 1

	if len(os.Args) == 5 {
		rounds, _ = strconv.Atoi(os.Args[4])
	}

	out := make([]rune, len(msg))
	copy(out, msg)

	if os.Args[1] == "encrypt" {
		out = vigenere_jorin.Encrypt(key, out, rounds)
	}

	if os.Args[1] == "decrypt" {
		out = vigenere_jorin.Decrypt(key, out, rounds)
	}

	fmt.Println(string(out))
}
