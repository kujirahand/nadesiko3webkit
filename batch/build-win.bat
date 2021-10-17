rem nadesiko3

set SCRIPT_DIR=%~dp0
cd %SCRIPT_DIR%
cd ..
set ROOT_DIR=%CD%\
set TARGET_DIR=%SCRIPT_DIR%win32

rem --- build ---
go build -ldflags="-H windowsgui" -o %TARGET_DIR%\nadesiko3.exe
xcopy %ROOT_DIR%webapp\* %TARGET_DIR%\webapp\ /s /d /e /h /r /y

rem --- icon ---

