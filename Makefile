# Makefile
all: setup build

setup:
	scripts/libcozo.sh

build:
	CGO_LDFLAGS="-L/${PWD}/libs" go build ./src/

test: 
	CGO_LDFLAGS="-L/${PWD}/libs" go test ./src/