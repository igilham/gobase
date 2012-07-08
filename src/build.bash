#!/bin/bash
# build all packages

OUTDIR="../_bin"

if [[ ! -d $OUTDIR ]]; then
  mkdir $OUTDIR
fi

go build -o $OUTDIR/basename basename/basename.go
go build -o $OUTDIR/cat cat/cat.go
go build -o $OUTDIR/cksum cksum/cksum.go
go build -o $OUTDIR/dirname dirname/dirname.go
go build -o $OUTDIR/echo echo/echo.go
go build -o $OUTDIR/false false/false.go
go build -o $OUTDIR/head head/head.go
go build -o $OUTDIR/ls ls/ls.go
go build -o $OUTDIR/mkdir mkdir/mkdir.go
go build -o $OUTDIR/pwd pwd/pwd.go
go build -o $OUTDIR/rm rm/rm.go
go build -o $OUTDIR/seq.exe seq/seq.go
go build -o $OUTDIR/sleep sleep/sleep.go
go build -o $OUTDIR/sort sort/sort.go
go build -o $OUTDIR/tee tee/tee.go
go build -o $OUTDIR/touch touch/touch.go
go build -o $OUTDIR/true true/true.go
go build -o $OUTDIR/wc wc/wc.go
go build -o $OUTDIR/whoami whoami/whoami.go
go build -o $OUTDIR/yes yes/yes.go
go build -o $OUTDIR/uniq uniq/uniq.go
