GOPATH=$(shell pwd)
export GOPATH

TARGETS=basename \
        cat \
        cksum \
				dirname \
				echo \
				false \
				head \
				ls \
				mkdir \
				pwd \
				rm \
				seq \
				sleep \
				sort \
				tee \
				touch \
				true \
				uniq \
				wc \
				whoami \
				yes

# default target
.PHONY: build

# build code
build:
	go build $(TARGETS)

# run unit tests
test:
	go test $(TARGETS)

# run unit tests with benchmarking enabled
bench:
	go test -bench . $(TARGETS)

# fetch the coverage tool
covcmd:
	go get code.google.com/p/go.tools/cmd/cover

# run unit tests with coverage analysis enabled
cov: covcmd
	go test -cover $(TARGETS)

# install gem dependencies for UAT test
gems:
	bundle install

# run cucumber user acceptance tests
uat: gems install
	cucumber

# automate formatting of code
fmt:
	go fmt $(TARGETS)

# automate fixing of source code for current version of go
fix:
	go fix $(TARGETS)

# fetch the vet command tool
vetcmd:
	go get code.google.com/p/go.tools/cmd/vet

vet: vetcmd
	go vet $(TARGETS)

# install binaries
install: test
	go install $(TARGETS)

all: install uat

clean:
	go clean $(TARGETS)
	rm -rf bin 2>/dev/null
	rm -rf pkg 2>/dev/null
