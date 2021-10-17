//go:build darwin || linux
// +build darwin linux

package main

import (
	"github.com/webview/webview"
)

func main() {
	debug := true

	go StartServer()
	w := webview.New(debug)
	defer w.Destroy()
	info := ReadIndexJson()
	w.SetTitle(info.Title)
	w.SetSize(info.Width, info.Height, webview.HintNone)
	w.Navigate(GetIndexURI())
	w.Run()
}
