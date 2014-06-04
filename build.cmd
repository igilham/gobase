@echo off
rem driver script to build/install/test/fix/vet all projects
set GOPATH=%~dp0
set TESTABLE_TARGETS=gobase sort
set ALL_TARGETS=basename cat cksum dirname echo false head ls mkdir pwd rm seq sleep tee touch true uniq wc whoami yes %TESTABLE_TARGETS%

go test %TESTABLE_TARGETS%
bundle install
cucumber
go install %ALL_TARGETS%
