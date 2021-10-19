package main

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/zserge/lorca"
)

// アプリ専用の保存フォルダを得る
func GetUserDir() string {
	// get home path
	home := os.Getenv("HOMEPATH")
	if runtime.GOOS != "windows" {
		home = os.Getenv("HOME")
	}
	appid := url.PathEscape(GlobalInfo.AppId)
	return filepath.Join(home, ".nadesiko3", appid)
}

func GetUserFilename(name string) string {
	name = url.PathEscape(name)
	path := filepath.Join(GetUserDir(), name)
	return path
}

func SaveUserFile(name string, value string) error {
	// 保存フォルダの確認
	dir := GetUserDir()
	if !Exists(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
	// 保存
	path := GetUserFilename(name)
	err := ioutil.WriteFile(path, []byte(value), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func LoadUserFile(name string) (string, error) {
	path := GetUserFilename(name)
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func EnumUserFiles() ([]string, error) {
	path := GetUserDir()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var paths []string = []string{}
	for _, f := range files {
		name, _ := url.PathUnescape(f.Name())
		paths = append(paths, name)
	}
	return paths, nil
}

func BindApi(ui lorca.UI) {
	// 関数をバインド (ただし、Promiseとなる)
	ui.Bind("nako3api_save", func(name string, value string) bool {
		err := SaveUserFile(name, value)
		return (err == nil)
	})
	ui.Bind("nako3api_load", func(name string) string {
		value, err := LoadUserFile(name)
		if err == nil {
			return value
		}
		return ""
	})
	ui.Bind("nako3api_files", func() []string {
		files, err := EnumUserFiles()
		if err == nil {
			return files
		}
		return []string{}
	})
}
