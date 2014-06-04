GOPATH=$(shell pwd)

ALL_TARGETS=cksum echo false ls mkdir pwd rm seq sleep tee true whoami yes $(TESTABLE_TARGETS)
TESTABLE_TARGETS=gobase basename cat dirname head sort touch uniq wc

# default target
.PHONY: build

# build code
build:
	go build $(ALL_TARGETS)

# run unit tests
test:
	go test $(TESTABLE_TARGETS)

gems:
	bundle install

# run cucumber user acceptance tests
uat: gems install
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
install: test
	go install $(ALL_TARGETS)

all: install uat

clean:
	go clean $(ALL_TARGETS)
	rm -rf bin 2>/dev/null
	rm -rf pkg 2>/dev/null
