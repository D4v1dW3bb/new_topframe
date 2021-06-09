package main

import (
	"log"

	"github.com/d4v1dw3bb/webview2"
	"github.com/lxn/win"
)

func main() {
	debug := true
	w := webview2.New(debug)
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview2.HintFixed)
	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	w.SetTransparentBackground(win.HWND(w.Window()), 50)
	w.Run()

}
