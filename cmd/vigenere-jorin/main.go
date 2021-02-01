package main

import (
	"bufio"
	"fmt"
	"nning.io/go/vigenere_jorin"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 && len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, `
%s <operation> <key> <message> [rounds]

  operation    "encrypt" or "decrypt"
  key          key for operation
  message      text for operation or "-" for reading from stdin
  rounds       times to repeat encryption, default 1

Key and message will be transformed to only upper case letters and space.

`, os.Args[0])
		os.Exit(1)
	}

	key := vigenere_jorin.Sanitize(os.Args[2])

	m := os.Args[3]
	if m == "-" {
		m = ""

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			m = m + scanner.Text()

			if m[len(m)-1] != '\n' {
				m = m + "\n"
			}
		}
	}

	msg := vigenere_jorin.Sanitize(m)

	rounds := 1

	if len(os.Args) == 5 {
		rounds, _ = strconv.Atoi(os.Args[4])
	}

	out := make([]rune, len(msg))
	copy(out, msg)

	switch os.Args[1][0] {
	case 'e':
		out = vigenere_jorin.Encrypt(key, out, rounds)
	case 'd':
		out = vigenere_jorin.Decrypt(key, out, rounds)
	}

	fmt.Println(string(out))
}
