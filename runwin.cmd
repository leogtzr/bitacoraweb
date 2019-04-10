@echo off
rem Leo Gutierrez R. <leogutierrezramirez@gmail.com>
cd ..\bitacora || (
    exit /b 1
)
start "" npm start

cd ..\bitacoraweb || (
    exit /b 1
)

echo "Building code ... "
go build && (
    
    echo Running binary

    ping google.com > nul
    ping google.com > nul
    ping google.com > nul

    .\bitacoraweb.exe
)
echo OK

exit /b 0