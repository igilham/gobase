@echo off
rem build all projects
set OUTDIR="_bin"

if exist %OUTDIR% goto build
mkdir %OUTDIR%

:build

go build -o %OUTDIR%\basename.exe src\basename\basename.go
go build -o %OUTDIR%\cat.exe src\cat\cat.go
go build -o %OUTDIR%\cksum.exe src\cksum\cksum.go
go build -o %OUTDIR%\dirname.exe src\dirname\dirname.go
go build -o %OUTDIR%\echo.exe src\echo\echo.go
go build -o %OUTDIR%\false.exe src\false\false.go
go build -o %OUTDIR%\head.exe src\head\head.go
go build -o %OUTDIR%\ls.exe src\ls\ls.go
go build -o %OUTDIR%\mkdir.exe src\mkdir\mkdir.go
go build -o %OUTDIR%\pwd.exe src\pwd\pwd.go
go build -o %OUTDIR%\rm.exe src\rm\rm.go
go build -o %OUTDIR%\sleep.exe src\sleep\sleep.go
go build -o %OUTDIR%\sort.exe src\sort\sort.go
go build -o %OUTDIR%\tee.exe src\tee\tee.go
go build -o %OUTDIR%\touch.exe src\touch\touch.go
go build -o %OUTDIR%\true.exe src\true\true.go
go build -o %OUTDIR%\wc.exe src\wc\wc.go
go build -o %OUTDIR%\whoami.exe src\whoami\whoami.go
go build -o %OUTDIR%\yes.exe src\yes\yes.go
go build -o %OUTDIR%\uniq.exe src\uniq\uniq.go
