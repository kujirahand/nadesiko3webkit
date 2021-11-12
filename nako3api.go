package main

import (
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

func Nako3api_save(name string, value string) bool {
	err := SaveUserFile(name, value)
	return (err == nil)
}

func Nako3api_load(name string) string {
	value, err := LoadUserFile(name)
	if err == nil {
		return value
	}
	return ""
}

func Nako3api_files() []string {
	files, err := EnumUserFiles()
	if err == nil {
		return files
	}
	return []string{}
}

func Nako3api_exec(args []string) string {
	var out []byte
	var err error
	switch len(args) {
	case 0:
		return ""
	case 1:
		out, err = exec.Command(args[0]).Output()
	case 2:
		out, err = exec.Command(args[0], args[1]).Output()
	case 3:
		out, err = exec.Command(args[0], args[1], args[2]).Output()
	case 4:
		out, err = exec.Command(args[0], args[1], args[2], args[3]).Output()
	case 5:
		out, err = exec.Command(args[0], args[1], args[2], args[3], args[4]).Output()
	case 6:
		out, err = exec.Command(args[0], args[1], args[2], args[3], args[4], args[5]).Output()
	}
	if err != nil {
		return ""
	}
	return string(out)
}

func Nako3api_getenv(key string) string {
	return os.Getenv(key)
}

func Nako3api_setenv(key string, val string) error {
	return os.Setenv(key, val)
}

func Nako3api_envlist() []string {
	return os.Environ()
}

/*
func BindApi(w lorca.UI) {
	// 関数をバインド (ただし、Promiseとなる)
	w.Bind("Nako3api_save", Nako3api_save)
	w.Bind("Nako3api_load", Nako3api_load)
	w.Bind("Nako3api_files", Nako3api_files)
	w.Bind("Nako3api_exec", Nako3api_exec)
	w.Bind("Nako3api_getenv", Nako3api_getenv)
	w.Bind("Nako3api_setenv", Nako3api_setenv)
	w.Bind("Nako3api_envlist", Nako3api_envlist)
}
*/
