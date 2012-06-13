#!/bin/bash

OUTDIR="_bin"

if [[ ! -d $OUTDIR ]]; then
  mkdir $OUTDIR
fi

go build -o $OUTDIR/basename src/basename/basename.go
go build -o $OUTDIR/cat src/cat/cat.go
go build -o $OUTDIR/cksum src/cksum/cksum.go
go build -o $OUTDIR/dirname src/dirname/dirname.go
go build -o $OUTDIR/echo src/echo/echo.go
go build -o $OUTDIR/false src/false/false.go
go build -o $OUTDIR/head src/head/head.go
go build -o $OUTDIR/ls src/ls/ls.go
go build -o $OUTDIR/mkdir src/mkdir/mkdir.go
go build -o $OUTDIR/pwd src/pwd/pwd.go
go build -o $OUTDIR/rm src/rm/rm.go
go build -o $OUTDIR/seq.exe src/seq/seq.go
go build -o $OUTDIR/sleep src/sleep/sleep.go
go build -o $OUTDIR/sort src/sort/sort.go
go build -o $OUTDIR/tee src/tee/tee.go
go build -o $OUTDIR/touch src/touch/touch.go
go build -o $OUTDIR/true src/true/true.go
go build -o $OUTDIR/wc src/wc/wc.go
go build -o $OUTDIR/whoami src/whoami/whoami.go
go build -o $OUTDIR/yes src/yes/yes.go
go build -o $OUTDIR/uniq src/uniq/uniq.go
