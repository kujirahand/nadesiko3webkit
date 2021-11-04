#!/bin/bash
APP_NAME=nadesiko3
SCRIPT_DIR=$(cd $(dirname $0); pwd)
DEST_PATH=$SCRIPT_DIR/linux-nako3-chrome
ROOT_DIR=$(cd $SCRIPT_DIR; cd ..; pwd)

echo "COPY TEMPLATE"
rm -f -r $DEST_PATH
mkdir -p $DEST_PATH
cp $SCRIPT_DIR/res/README.md $DEST_PATH/

echo "COPY VIEW"
cp $ROOT_DIR/_view_lorca.go $ROOT_DIR/view.go

echo "BUILD"
cd $ROOT_DIR
GOOS=linux GOARCH=arm go build -o $DEST_PATH/$APP_NAME-arm32
GOOS=linux GOARCH=arm64 go build -o $DEST_PATH/$APP_NAME-arm64
GOOS=linux GOARCH=386 go build -o $DEST_PATH/$APP_NAME-386
GOOS=linux GOARCH=amd64 go build -o $DEST_PATH/$APP_NAME-amd64

echo "COPY RESOURCES"
cp -r $ROOT_DIR/webapp $DEST_PATH/webapp
echo "done."


