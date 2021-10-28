# nadesiko3webkit

これは、[日本語プログラミング言語「なでしこ3(Web版)」](https://nadesi.com/)で作ったプログラムを配布したい場合に便利なパッケージです。
なでしこで作ったゲームやツールを一般配布するのにご利用ください。

ソースからビルドする場合は、Go言語が必要になりますが、なでしこ3(Web版)のプログラムを配布したいだけであれば、[こちら](https://github.com/kujirahand/nadesiko3webkit/releases)から配布用パッケージをダウンロードするだけでOKです。

なでしこ3で作ったプログラムを配布するには、Electronを使う方法もありますが、本プロジェクト(nadesiko3webkit)を使うと、OSにインストールしたChromeのコンポーネントを使ってなでしこ3を実行するので、配布サイズが小さく手軽にプログラムを配布できるというメリットがあります。簡単なプログラムであれば、ZIP圧縮して5MB程度の配布サイズになります。

詳しい使い方は、[こちら](batch/res/README.md)をご覧ください。

## 本ライブラリのコンパイル

本ライブラリを構築するには、Go言語が必要です。

 - [Go言語](https://golang.org/)
 - 利用パッケージ (以下を選択してビルド)
   - [lorca](https://github.com/zserge/lorca)
   - [webview](https://github.com/webview/webview)
   - [webview2](https://github.com/jchv/go-webview2)

## コンパイルの方法

必要なモジュールを取得します。

```bash
go get -u
```

## コンポーネントの切り替え

以下のいずれかのバッチを実行します。

```bash
# lorcaを使う場合(win/mac)
./make_lorca.sh
.\make_lorca.bat

# WebViewを使う場合(mac)
./make_webview_mac.sh

# WebViewを使う場合(win)
./make_webview_win.bat
```

## 配布ファイルの作成

batchフォルダ以下のバッチを実行します。

### Windowsの場合

```bash
batch¥build-win.bat
```

### macOSの場合

```bash
batch/build-mac.sh
```

