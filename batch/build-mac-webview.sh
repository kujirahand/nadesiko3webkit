#!/bin/bash
APP_NAME=nadesiko3

SCRIPT_DIR=$(cd $(dirname $0); pwd)
MACOS_PATH=$SCRIPT_DIR/mac-nako3-webview
APP_PATH=$MACOS_PATH/$APP_NAME.app
TEMPLATE_PATH=$SCRIPT_DIR/res/Template.app
ROOT_DIR=$(cd $SCRIPT_DIR; cd ..; pwd)

echo "COPY TEMPLATE"
rm -f -r $MACOS_PATH
mkdir -p $MACOS_PATH
cp -r $TEMPLATE_PATH $APP_PATH
mkdir -p $APP_PATH/Contents/Resources
cp $SCRIPT_DIR/res/README.md $MACOS_PATH/

echo "COPY VIEW"
cp $ROOT_DIR/_view_webview_mac.go $ROOT_DIR/view.go

echo "BUILD"
cd $ROOT_DIR
### (memo) 
### go version 1.17.2 では、M1 Macを使う時、HomebrewでgoをIntel版とM1版を両方にインストールしておく必要があった
### WebViewのリンクができなかった
GO_M="/opt/homebrew/bin/go"
GO_I="/usr/local/bin/go"
C_MAC_DIR=$APP_PATH/Contents/MacOS
APP_FULL=$C_MAC_DIR/$APP_NAME
APP_ARM=$C_MAC_DIR/$APP_NAME.arm64
APP_AMD=$C_MAC_DIR/$APP_NAME.amd64
mkdir -p $C_MAC_DIR
if [[ -e $GO_M ]]; then
  #   echo "[OK] Build mac universal binary"
  #   # MAKE Universal Binary 
  #   # (ref) https://dev.to/thewraven/universal-macos-binaries-with-go-1-16-3mm3
  #   GOOS=darwin GOARCH=arm64 $GO_M build -o $APP_ARM
  #   GOOS=darwin GOARCH=amd64 $GO_M build -o $APP_AMD
  #   # なぜか出力ディレクトリにフルパスは指定できない仕様のようだ
  #   cd $C_MAC_DIR
  #   lipo -create -output $APP_NAME $APP_ARM $APP_AMD
  #   file $APP_FULL
  #   rm $APP_ARM $APP_AMD
  echo "Build M1 mac binary"
  GOOS=darwin GOARCH=arm64 $GO_M build -o $APP_FULL
else
  echo "[Warning] build binary NOT universal"
  go build -o $APP_FULL
fi

echo "COPY RESOURCES"
cp -r $ROOT_DIR/webapp $APP_PATH/Contents/MacOS/webapp

echo "done."



