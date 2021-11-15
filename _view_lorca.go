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

func BindApi(w lorca.UI) {
	// 関数をバインド (ただし、Promiseとなる)
	w.Bind("Nako3api_info", Nako3api_info)
	w.Bind("Nako3api_save", Nako3api_save)
	w.Bind("Nako3api_load", Nako3api_load)
	w.Bind("Nako3api_files", Nako3api_files)
	w.Bind("Nako3api_exec", Nako3api_exec)
	w.Bind("Nako3api_getenv", Nako3api_getenv)
	w.Bind("Nako3api_setenv", Nako3api_setenv)
	w.Bind("Nako3api_envlist", Nako3api_envlist)
}
