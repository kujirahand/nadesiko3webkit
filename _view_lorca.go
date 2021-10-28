package main

import (
	"log"

	"github.com/zserge/lorca"
)

func ShowBrowser(info *IndexInfo) {
	// Check Chrome runtime
	if lorca.LocateChrome() == "" {
		lorca.PromptDownload()
		log.Fatal("cannot find chrome app")
	}

	// ブラウザを起動
	indexUrl := GetIndexPageURL()
	ui, err := lorca.New(indexUrl, "", info.Width, info.Height)
	if err != nil {
		log.Fatal("cannot open browser")
	}
	defer ui.Close()

	// 独自APIを登録
	BindApi(ui)

	<-ui.Done()
}

func BindApi(ui lorca.UI) {
	// 関数をバインド (ただし、Promiseとなる)
	ui.Bind("nako3api_save", Nako3api_save)
	ui.Bind("nako3api_load", Nako3api_load)
	ui.Bind("nako3api_files", Nako3api_files)
}
