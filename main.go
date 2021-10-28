package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"
)

// index.json
type IndexInfo struct {
	Title  string `json:"title"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Port   int    `json:"port"`
	AppId  string `json:"appid"`
	Resize bool   `json:"resize"`
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

func GetIndexPageURL() string {
	utime := strconv.FormatInt(time.Now().Unix(), 16)
	return "http://127.0.0.1:" + strconv.Itoa(GlobalInfo.Port) + "/webapp/index.html?time=" + utime
}

func main() {
	// Load Setting file (index.json)
	info := ReadIndexJson()

	// ローカルサーバーを起動
	go StartServer(&info)

	// ブラウザを起動
	ShowBrowser(&info)
}
