#!/bin/bash

APP_NAME=nadesiko3

SCRIPT_DIR=$(cd $(dirname $0); pwd)
TARGET_PATH=$SCRIPT_DIR/$APP_NAME.app
TEMPLATE_PATH=$SCRIPT_DIR/res/Template.app

cd $SCRIPT_DIR
cd ..
rm -f -r $TARGET_PATH
cp -r $TEMPLATE_PATH $TARGET_PATH
mkdir -p $TARGET_PATH/Contents/MacOS
mkdir -p $TARGET_PATH/Contents/Resources

go build -o $TARGET_PATH/Contents/MacOS/$APP_NAME
cp -r ./webapp $TARGET_PATH/Contents/MacOS/



