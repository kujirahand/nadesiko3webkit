#!/bin/bash

APP_NAME=nadesiko3

SCRIPT_DIR=$(cd $(dirname $0); pwd)
MACOS_PATH=$SCRIPT_DIR/mac-nako3
TARGET_PATH=$MACOS_PATH/$APP_NAME.app
TEMPLATE_PATH=$SCRIPT_DIR/res/Template.app

cd $SCRIPT_DIR
cd ..
rm -f -r $MACOS_PATH

echo "COPY TEMPLATE"
mkdir -p $MACOS_PATH
cp -r $TEMPLATE_PATH $TARGET_PATH
mkdir -p $TARGET_PATH/Contents/Resources
cp $SCRIPT_DIR/res/README.md $MACOS_PATH/

echo "BUILD"
go build -o $TARGET_PATH/Contents/MacOS/$APP_NAME
echo "COPY RESOURCES"
cp -r ./webapp $TARGET_PATH/Contents/MacOS/



