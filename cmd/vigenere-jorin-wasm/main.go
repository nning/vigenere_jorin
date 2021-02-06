package main

import (
	"syscall/js"

	. "nning.io/go/vigenere_jorin"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("encrypt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := Sanitize(args[0].String())
		msg := Sanitize(args[1].String())

		return string(Encrypt(key, msg, 1))
	}))

	js.Global().Set("decrypt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := Sanitize(args[0].String())
		msg := Sanitize(args[1].String())

		return string(Decrypt(key, msg, 1))
	}))

	<-c
}
