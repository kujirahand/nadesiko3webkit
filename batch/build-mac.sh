SCRIPT_DIR=$(cd $(dirname $0); pwd)
APP_NAME=nadesiko3

cd $SCRIPT_DIR
cd ..
mkdir -p $APP_NAME.app/Contents/MacOS
go build -o $APP_NAME.app/Contents/MacOS/$APP_NAME
cp -r ./webapp $APP_NAME.app/Contents/MacOS/



