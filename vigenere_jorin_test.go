package vigenere_jorin

import (
	"testing"
)

func TestInversibility(t *testing.T) {
	key := Sanitize("joocuwu ieviecoh")
	msg := Sanitize("pheixiegeis eshug ochakohzo ee iecheeghievo queedei ooxa")

	c := Encrypt(key, msg)
	p := Decrypt(key, c)

	if string(p) != string(msg) {
		t.Error("Non-inversible")
	}
}
