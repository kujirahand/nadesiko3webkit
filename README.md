# nadesiko3webkit

これは、[日本語プログラミング言語「なでしこ3(Web版)」](https://nadesi.com/)で作ったプログラムを配布したい場合に便利なパッケージです。
なでしこで作ったゲームやツールを一般配布するのにご利用ください。

ソースからビルドする場合は、Go言語が必要になりますが、なでしこ3(Web版)のプログラムを配布したいだけであれば、[こちら](https://github.com/kujirahand/nadesiko3webkit/releases)から配布用パッケージをダウンロードするだけでOKです。

なでしこ3で作ったプログラムを配布するには、Electronを使う方法もありますが、本プロジェクト(nadesiko3webkit)を使うと、OSにインストールされているWebKitベースのブラウザのコンポーネントを使ってなでしこ3を実行するので、配布サイズが小さく手軽にプログラムを配布できるというメリットがあります。簡単なプログラムであれば、ZIP圧縮して5MB程度の配布サイズになります。

詳しい使い方は、[こちら](res/../README.md)をご覧ください。

# 本ライブラリのコンパイル

本ライブラリを構築するには、Go言語が必要です。

 - [Go言語](https://golang.org/)
 - 利用パッケージ
   - [lorca](https://github.com/zserge/lorca)

## コンパイルの方法

必要なモジュールを取得します。

```bash
go get -u
```

## Windowsの場合

以下のコマンドを実行します。

```bash
batch¥build-win.bat
```

## macOSの場合

以下のコマンドを実行します。

```bash
batch/build-mac.sh
```

