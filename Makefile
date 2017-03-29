DIST := dist

ifeq ($(OS), Windows_NT)
	EXECUTABLE := dbweb.exe
else
	EXECUTABLE := dbweb
endif

BINDATA := modules/{options,static,templates}/bindata.go

LDFLAGS := -X "main.Version=$(shell git describe --tags --always | sed 's/-/+/' | sed 's/^v//')" -X "main.Tags=$(TAGS)"

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f)

TAGS ?=

ifneq ($(DRONE_TAG),)
	VERSION ?= $(subst v,,$(DRONE_TAG))
else
	ifneq ($(DRONE_BRANCH),)
		VERSION ?= $(subst release/v,,$(DRONE_BRANCH))
	else
		VERSION ?= master
	endif
endif

.PHONY: all
all: build

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf $(EXECUTABLE) $(DIST) $(BINDATA)

.PHONY: generate
generate:
	@hash go-bindata > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/jteeuwen/go-bindata/...; \
	fi
	go generate $(PACKAGES)

.PHONY: build
build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	go build -i -v -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -o $@