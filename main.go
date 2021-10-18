package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"github.com/zserge/lorca"
)

// index.json
type IndexInfo struct {
	Title  string `json:"title"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Port   int    `json:"port"`
	AppId  string `json:"appid"`
}

var GlobalInfo *IndexInfo

const DIR_WEBAPP = "webapp"

func ReadIndexJson() IndexInfo {
	// index.json
	var info IndexInfo = IndexInfo{Title: "なでしこ3", Width: 800, Height: 600, Port: 8888}
	indexJson := filepath.Join(GetBokanPath(), "webapp", "index.json")
	raw, err := ioutil.ReadFile(indexJson)
	if err == nil { // 読み込めた時
		json.Unmarshal(raw, &info)
		fmt.Printf("size=%d,%d\n", info.Width, info.Height)
	}
	checkPort(&info)
	GlobalInfo = &info
	return info
}

func GetIndexURI(info *IndexInfo) string {
	utime := strconv.FormatInt(time.Now().Unix(), 16)
	return "http://127.0.0.1:" + strconv.Itoa(info.Port) + "/webapp/index.html?time=" + utime
}

func main() {
	// Check Chrome runtime
	if lorca.LocateChrome() == "" {
		lorca.PromptDownload()
		log.Fatal(fmt.Errorf("cannot open Chrome app"))
	}

	// Load Setting file (index.json)
	info := ReadIndexJson()

	// ローカルサーバーを起動
	go StartServer(&info)

	// ブラウザを起動
	ui, _ := lorca.New(GetIndexURI(&info), "", info.Width, info.Height)
	defer ui.Close()

	// 関数をバインド (ただし、Promiseとなる)
	ui.Bind("nako3api_save", func(name string, value string) bool {
		err := SaveData(name, value)
		return (err == nil)
	})
	ui.Bind("nako3api_load", func(name string) string {
		value, err := LoadData(name)
		if err == nil {
			return value
		}
		return ""
	})
	ui.Bind("nako3api_files", func() []string {
		files, err := GetFiles()
		if err == nil {
			return files
		}
		return []string{}
	})
	<-ui.Done()
}
