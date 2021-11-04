#!/bin/bash
APP_NAME=nadesiko3
SCRIPT_DIR=$(cd $(dirname $0); pwd)
MAC_PATH=$SCRIPT_DIR/mac-nako3-chrome
APP_PATH=$MAC_PATH/$APP_NAME.app
TEMPLATE_PATH=$SCRIPT_DIR/res/Template.app
ROOT_DIR=$(cd $SCRIPT_DIR; cd ..; pwd)


echo "COPY TEMPLATE"
rm -f -r $MAC_PATH
mkdir -p $MAC_PATH
cp -r $TEMPLATE_PATH $APP_PATH
mkdir -p $APP_PATH/Contents/Resources
cp $SCRIPT_DIR/res/README.md $MAC_PATH/

echo "COPY VIEW"
cp $ROOT_DIR/_view_lorca.go $ROOT_DIR/view.go

echo "BUILD"
C_MAC_DIR=$APP_PATH/Contents/MacOS
ARM=$C_MAC_DIR/$APP_NAME.arm64
AMD=$C_MAC_DIR/$APP_NAME.amd64
cd $ROOT_DIR
GOOS=darwin GOARCH=amd64 go build -o $ARM
GOOS=darwin GOARCH=arm64 go build -o $AMD
mkdir -p $C_MAC_DIR
cd $C_MAC_DIR
lipo -create -output $APP_NAME $ARM $AMD
rm $ARM $AMD

echo "COPY RESOURCES"
cp -r $ROOT_DIR/webapp $APP_PATH/Contents/MacOS/webapp
echo "done."


