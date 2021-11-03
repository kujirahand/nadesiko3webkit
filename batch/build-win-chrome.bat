@echo off
@rem nadesiko3

set SCRIPT_DIR=%~dp0
cd %SCRIPT_DIR%
cd ..
set ROOT_DIR=%CD%\
set TARGET_DIR=%SCRIPT_DIR%win-nako3-chrome\

echo --------------------------------------------------------------
echo ROOT_DIR=%ROOT_DIR%
echo SCRIPT_DIR=%SCRIPT_DIR%
echo TARGET_DIR=%TARGET_DIR%
echo --------------------------------------------------------------

rem --- prepare webview2 ---
copy %ROOT_DIR%_view_lorca.go %ROOT_DIR%view.go

rem --- reset ---
rmdir /s /q %TARGET_DIR%
mkdir %TARGET_DIR%

rem --- build ---
copy %SCRIPT_DIR%res\nadesiko3webkit.syso %ROOT_DIR%nadesiko3webkit.syso
go build -ldflags="-H windowsgui" -o %TARGET_DIR%nadesiko3.exe
xcopy %ROOT_DIR%webapp\* %TARGET_DIR%webapp\ /s /d /e /h /r /y
copy %SCRIPT_DIR%res\README.md %TARGET_DIR%\

echo "ok"
pause

