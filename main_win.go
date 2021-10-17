//go:build windows
// +build windows

package main

import (
	webview "github.com/jchv/go-webview2"
)

func main() {
	debug := true
	info := ReadIndexJson()

	// ローカルサーバーを起動
	go StartServer(info.Port)

	// ブラウザを起動
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle(info.Title)
	hint := webview.HintNone
	if !info.Resize {
		hint = webview.HintFixed
	}
	w.SetSize(info.Width, info.Height, hint)
	w.Navigate(GetIndexURI(info.Port))
	w.Run()
}
