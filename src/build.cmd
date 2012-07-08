@echo off
rem build all projects
set OUTDIR="..\_bin"

if exist %OUTDIR% goto build
mkdir %OUTDIR%

:build

go build -o %OUTDIR%\basename.exe basename\basename.go
go build -o %OUTDIR%\cat.exe cat\cat.go
go build -o %OUTDIR%\cksum.exe cksum\cksum.go
go build -o %OUTDIR%\dirname.exe dirname\dirname.go
go build -o %OUTDIR%\echo.exe echo\echo.go
go build -o %OUTDIR%\false.exe false\false.go
go build -o %OUTDIR%\head.exe head\head.go
go build -o %OUTDIR%\ls.exe ls\ls.go
go build -o %OUTDIR%\mkdir.exe mkdir\mkdir.go
go build -o %OUTDIR%\pwd.exe pwd\pwd.go
go build -o %OUTDIR%\rm.exe rm\rm.go
go build -o %OUTDIR%\seq.exe seq\seq.go
go build -o %OUTDIR%\sleep.exe sleep\sleep.go
go build -o %OUTDIR%\sort.exe sort\sort.go
go build -o %OUTDIR%\tee.exe tee\tee.go
go build -o %OUTDIR%\touch.exe touch\touch.go
go build -o %OUTDIR%\true.exe true\true.go
go build -o %OUTDIR%\wc.exe wc\wc.go
go build -o %OUTDIR%\whoami.exe whoami\whoami.go
go build -o %OUTDIR%\yes.exe yes\yes.go
go build -o %OUTDIR%\uniq.exe uniq\uniq.go
