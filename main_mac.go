//go:build darwin || linux
// +build darwin linux

package main

import (
	"github.com/webview/webview"
)

func main() {
	debug := true
	info := ReadIndexJson()

	// ローカルサーバーを起動
	go StartServer(&info)

	// ブラウザを起動
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle(info.Title)
	hint := webview.HintNone
	if !info.Resize {
		hint = webview.HintFixed
	}
	w.SetSize(info.Width, info.Height, webview.Hint(hint))
	w.Navigate(GetIndexURI(&info))
	w.Run()
}
