# なでしこ3配布用Win/Macパッケージ

このアーカイブは、日本語プログラミング言語「なでしこ3(Web版)」で作ったプログラムを実行ファイルとしてユーザーに配布したい場合に便利なツールです。

Windows版とmacOS版のバイナリファイルが提供されています。

 - [配布URL] https://github.com/kujirahand/nadesiko3webkit/releases

# ダウンロードと動作確認

上記の配布URLから、win-nako3.zip / mac-nako3.zip をダウンロードします。
ZIPファイルを解凍したら、nadesiko3の実行ファイルをダブルクリックしてみてください。
無事になでしこ3が動くごとが分かったら次の作業に移りましょう。

なお、本アプリを動かすには、Chrome(またはChromium)が必要です。Chromeのコンポーネントを利用して動作します。
以下よりインストールしてください。

 - [Google Chrome](https://www.google.com/intl/ja_jp/chrome/)
   - 要求バージョン: Chrome/Chromium >= 70

# プログラムを差し替えよう

配布用プログラムには、webappフォルダが含まれています。このフォルダを開くと「main.nako3」というファイルがあります。
このファイルが起動してすぐに実行されるプログラムです。
ウィンドウサイズやタイトルを変更したい場合は「index.json」を編集してください。

macOSの場合は、nadesiko3.app/Contents/MacOS/webappにwebappディレクトリがあります。
nadesiko3.appを右クリックして「パッケージ内容を表示」を選択すると、Finderにファイルが表示されます。

## プログラムがうまく動かない場合

プログラムのエラーで動かない場合、画面を右クリックして「検証」を選択すると、デベロッパーツールが表示されます。
「Console」タブを開いてエラーが出ていないか確認してみましょう。
できるだけ、ユーザーにエラーを見せないように配慮して、わざとエラーを画面に出さないようにしています。

# 特殊コマンドについて

特別なコマンドが利用できます。
ファイルの保存と読み込みが可能です。いずれも非同期に実行されます。

```api-list.nako3
# ファイルへの保存
「test.txt」に「あいうえお」をファイル保存した時には
　　「保存しました」と表示
ここまで。

# ファイルの読み込み
「test.txt」をファイル読んだ時には
　　対象を表示。
ここまで。

# ファイル一覧の取得
ファイル一覧取得した時には
　　対象を反復
　　　　それを表示。
　　ここまで。
ここまで。

# コマンドを実行して結果を得る
# Windows
「calc」を起動した時には
　　対象を表示。
ここまで。
# macOSで電卓アプリを起動する
["open", "/System/Applications/Calculator.app"]を起動した時には
　　対象を表示。
ここまで。
# macOSでlsコマンドを実行して結果を得る
["ls", "-a"]を起動時には
　　対象を表示。
ここまで。

# 環境変数の取得と設定
「HOME」の環境変数取得した時には
　　対象を表示。
ここまで。
環境変数一覧取得時には
　　対象をJSONエンコード整形して表示。
ここまで。
「VAR1」へ「hoge」を環境変数設定した時には(ERR)
　　# 設定後の処理をここに
ここまで。
```

セキュリティの問題を考慮して、ファイルを任意のフォルダに保存することはできません。
ファイルは以下のフォルダに保存されます。なお、index.jsonでappidを指定します。

```text
<ユーザーフォルダ>/.nadesiko3/<appid>
```

## 特殊コマンド利用例

以下ファイルを保存して読み込む利用例です。

```test.nako3
「いろはにほへと」を「test.txt」へファイル保存時には
　　「保存しました」と言う。
　　「test.txt」からファイル読時には
　　　　対象を言う。
　　ここまで。
ここまで。
```
