#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)
rm -f $SCRIPT_DIR/webapp/nadesiko3/*
cd ../nadesiko3
npm run build
cp ./release/* $SCRIPT_DIR/webapp/nadesiko3/
