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
	var defaultChromeArgs = []string{
		"--disable-background-networking",
		"--disable-background-timer-throttling",
		"--disable-backgrounding-occluded-windows",
		"--disable-breakpad",
		"--disable-client-side-phishing-detection",
		"--disable-default-apps",
		"--disable-dev-shm-usage",
		"--disable-infobars",
		"--disable-extensions",
		"--disable-features=site-per-process",
		"--disable-hang-monitor",
		"--disable-ipc-flooding-protection",
		"--disable-popup-blocking",
		"--disable-prompt-on-repost",
		"--disable-renderer-backgrounding",
		"--disable-sync",
		"--disable-translate",
		"--disable-windows10-custom-titlebar",
		"--metrics-recording-only",
		//"--no-first-run",//←こちらは元の記述
		"no-first-run", //←このように「--」を消す
		//"--start-fullscreen",             //  <- 起動時最大化
		"--no-default-browser-check",
		"--safebrowsing-disable-auto-update",
		"--disable-automation", // 自動テスト ソフトウェアによって制御されています」を消したい(実際は消えない)
		"--password-store=basic",
		"--use-mock-keychain",
		"--remote-allow-origins=*", //  <- websoket.Dialのbad statusエラーを出さない
	}
	indexUrl := GetIndexPageURL()
	ui, err := lorca.New(indexUrl, "", info.Width, info.Height, defaultChromeArgs...)
	if err != nil {
		log.Fatal(err)
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
