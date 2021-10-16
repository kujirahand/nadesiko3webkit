rem nadesiko3
cd %~dp0
cd ..
go build -ldflags="-H windowsgui" -o bin\nadesiko3.exe
xcopy .\webapp bin\webapp /d /e /h /r /y
explorer .\bin
bin\nadesiko3.exe
pause
