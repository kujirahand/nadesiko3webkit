//go:build windows
// +build windows

package main

import (
	webview "github.com/jchv/go-webview2"
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
