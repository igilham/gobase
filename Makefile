GOPATH=$(shell pwd)

ALL_TARGETS=cksum dirname echo false head ls mkdir pwd rm seq sleep tee touch true uniq wc whoami yes $(TESTABLE_TARGETS)
TESTABLE_TARGETS=gobase basename cat sort

# default target
.PHONY: build

# build code
build:
	go build $(ALL_TARGETS)

# run all tests
test: unittest cucumber

# run unit tests
unittest:
	go test $(TESTABLE_TARGETS)

gems:
	bundle install

# run cukes
cucumber: gems
	cucumber

# automate formatting of code
fmt:
	go fmt $(ALL_TARGETS)

# automate fixing of source code for current version of go
fix:
	go fix $(ALL_TARGETS)

# fetch the vet command tool
vetcmd:
	go get code.google.com/p/go.tools/cmd/vet

vet: vetdcmd
	go vet $(ALL_TARGETS)

# install binaries
install:
	go install $(ALL_TARGETS)

clean:
	go clean $(ALL_TARGETS)
	rm -rf bin 2>/dev/null
	rm -rf pkg 2>/dev/null
