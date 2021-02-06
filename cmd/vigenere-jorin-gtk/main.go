package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"nning.io/go/vigenere_jorin"
)

const appID = "io.nning.vigenere-jorin.ui"

func main() {
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	panicOnError(err)

	application.Connect("activate", func() {
		builder, _ := gtk.BuilderNewFromFile("main.glade")

		buttonObj, _ := builder.GetObject("Button 1")
		button := buttonObj.(*gtk.Button)

		toggleObj, _ := builder.GetObject("Toggle")
		toggle := toggleObj.(*gtk.Switch)

		signals := map[string]interface{}{
			"start": func() {
				keyEntryObj, _ := builder.GetObject("Entry Key")
				keyEntry := keyEntryObj.(*gtk.Entry)

				key, _ := keyEntry.GetText()

				textEntryObj, _ := builder.GetObject("Entry Text")
				textEntry := textEntryObj.(*gtk.Entry)

				text, _ := textEntry.GetText()

				if key != "" && text != "" {
					sKey := vigenere_jorin.Sanitize(key)
					sText := vigenere_jorin.Sanitize(text)

					out := make([]rune, len(sText))
					copy(out, sText)

					if toggle.GetActive() {
						out = vigenere_jorin.Encrypt(sKey, out, 1)
					} else {
						out = vigenere_jorin.Decrypt(sKey, out, 1)
					}

					resultEntryObj, _ := builder.GetObject("Entry Result")
					resultEntry := resultEntryObj.(*gtk.Entry)

					resultEntry.SetText(string(out))
				}
			},
			"toggle": func() {
				if toggle.GetActive() {
					button.SetLabel("Encrypt")
				} else {
					button.SetLabel("Decrypt")
				}
			},
		}

		builder.ConnectSignals(signals)

		winObj, _ := builder.GetObject("Window Main")
		win := winObj.(*gtk.Window)

		win.SetTitle("Vigenere-Jorin")

		win.Show()
		application.AddWindow(win)
	})

	os.Exit(application.Run(os.Args))
}

func panicOnError(e error) {
	if e != nil {
		log.Panic(e)
	}
}
