export GO111MODULE=on
export GOOS=linux

.DEFAULT_GOAL := build

.PHONY: build
build: 
	GOOS=${GOOS} GO111MODULE=${GO111MODULE} go build -o=dist/rrs -v ${LDFLAGS} ${GCFLAGS} github.com/Lachann/rrs/cmd
