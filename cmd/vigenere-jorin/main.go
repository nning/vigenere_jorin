package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
	"nning.io/go/vigenere_jorin"
)

// Mode represents a mode object in the configuration
type Mode struct {
	Name             string `yaml:"name"`
	Alphabet         string `yaml:"alphabet"`
	KeyPositionReset bool   `yaml:"keyPositionReset"`
}

// Conf represents the top level object in the configuration
type Conf struct {
	DefaultMode string `yaml:"defaultMode"`
	Modes       []Mode `yaml:"modes"`
}

var mode = flag.String("m", "default", "Select mode defined in config.yml")

func printHelp(config *Conf) {
	fmt.Fprintf(os.Stderr, `
%s [options] <operation> <key> <message> [rounds]

	operation    "encrypt" or "decrypt"
	key          Key for operation
	message      Text for operation or "-" for reading from stdin
	rounds       Times to repeat operation, default 1

	options      Optional flags
		-m         Mode (from config.yml)

Key and message will be transformed to only upper case letters and space.

`, os.Args[0])

	if config != nil {
		fmt.Fprintf(os.Stderr, "Available modes: %v\n\n", getModeNames(config))
	}

	os.Exit(1)
}

func getConfig() *Conf {
	var c Conf

	content, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(content, &c)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	return &c
}

func getDefaultMode(config *Conf) string {
	if config.DefaultMode != "" {
		return config.DefaultMode
	}

	return "default"
}

func getMode(config *Conf, name string) *Mode {
	i := 0

	for ; i < len(config.Modes); i++ {
		if config.Modes[i].Name == name {
			break
		}
	}

	if i < len(config.Modes) {
		return &config.Modes[i]
	}

	return nil
}

func getModeNames(config *Conf) []string {
	a := make([]string, len(config.Modes))

	for i := 0; i < len(config.Modes); i++ {
		a[i] = config.Modes[i].Name
	}

	return a
}

func main() {
	args := os.Args
	argsLen := len(args)

	config := getConfig()

	if argsLen < 4 || argsLen > 7 {
		printHelp(config)
	}

	flag.Parse()

	args = flag.Args()
	argsLen = len(args)

	if argsLen < 3 || argsLen > 4 {
		printHelp(config)
	}

	if config != nil {
		defaultMode := getDefaultMode(config)
		mode = &defaultMode

		cMode := getMode(config, *mode)

		if mode != nil {
			vigenere_jorin.SetParameters(cMode.Alphabet, cMode.KeyPositionReset)
		} else {
			fmt.Fprintf(os.Stderr, "Warning: Mode definition for \"%s\" not found\n", *mode)
		}
	}

	key := vigenere_jorin.Sanitize(args[1])

	m := args[2]
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

	if argsLen == 4 {
		rounds, _ = strconv.Atoi(args[3])
	}

	out := make([]rune, len(msg))
	copy(out, msg)

	switch args[0][0] {
	case 'e':
		out = vigenere_jorin.Encrypt(key, out, rounds)
	case 'd':
		out = vigenere_jorin.Decrypt(key, out, rounds)
	}

	fmt.Println(string(out))
}
