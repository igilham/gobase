@echo off
rem driver script to build/install/test/fix/vet all projects
set GOPATH=%~dp0
set LIBS=gobase
set APPS=basename cat cksum dirname echo false head ls mkdir pwd rm seq sleep sort tee touch true uniq wc whoami yes

go %COMMAND% %LIBS% %APPS%

