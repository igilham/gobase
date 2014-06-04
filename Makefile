GOPATH=$(shell pwd)

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

vet: vetdcmd
	go vet $(TARGETS)

# install binaries
install: test
	go install $(TARGETS)

all: install uat

clean:
	go clean $(TARGETS)
	rm -rf bin 2>/dev/null
	rm -rf pkg 2>/dev/null
