package main

import (
	webview "github.com/jchv/go-webview2"
)

func ShowBrowser(info *IndexInfo) {

	// ブラウザを起動
	indexUrl := GetIndexPageURL()

	// ブラウザを起動
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle(info.Title)
	hint := webview.HintNone
	if !info.Resize {
		hint = webview.HintFixed
	}
	w.SetSize(info.Width, info.Height, webview.Hint(hint))
	// 独自APIを登録
	BindApi(w)
	w.Navigate(indexUrl)
	w.Run()
}

func BindApi(w webview.WebView) {
	// 関数をバインド (ただし、Promiseとなる)
	w.Bind("nako3api_save", Nako3api_save)
	w.Bind("nako3api_load", Nako3api_load)
	w.Bind("nako3api_files", Nako3api_files)
	w.Bind("nako3api_exec", Nako3api_exec)
}
