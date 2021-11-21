TARGET := go-commitizen
GOFILES := $(wildcard *.go) $(wildcard commitizen/*.go) $(wildcard git/*.go) $(wildcard model/*.go) $(wildcard prompt/*.go)

ifeq ($(OS),Windows_NT)
	GOOS := windows
	COPY := copy
else
	COPY := cp
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		GOOS := linux
	else ifeq ($(UNAME_S),Darwin)
		GOOS := darwin
	endif
endif

GIT_EXEC_PATH := $(shell git --exec-path)

all: ${TARGET}
install:
	$(COPY) go-commitizen $(GIT_EXEC_PATH)/git-cz
clean:
	rm -rf ${TARGET}


go-commitizen: $(GOFILES)
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -o $@

.PHONY: all install clean