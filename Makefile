SHELL=/usr/bin/env bash

project_name=computing-provider

unexport GOFLAGS

GOCC?=go

commitflag=

ldflags=-X=github.com/swanchain/go-computing-provider/build.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))+ $(commitflag)


all: mainnet
.PHONY: all

check_git_status:
	@if ! git diff --quiet; then \
		commitflag=dirty; \
	fi
	$(eval commitflag=$(commitflag))


computing-provider: check_git_status
computing-provider:
	rm -rf computing-provider
	$(GOCC) build $(GOFLAGS) -o computing-provider ./cmd/computing-provider
.PHONY: computing-provider

install:
	sudo install -C computing-provider /usr/local/bin/computing-provider

clean:
	sudo rm -rf /usr/local/bin/computing-provider
.PHONY: clean

mainnet: GOFLAGS+= -ldflags="$(ldflags) -X github.com/swanchain/go-computing-provider/build.NetWorkTag=mainnet"
mainnet: computing-provider

testnet: GOFLAGS+= -ldflags="$(ldflags) -X github.com/swanchain/go-computing-provider/build.NetWorkTag=proxima"
testnet: computing-provider

dev: GOFLAGS+= -ldflags="$(ldflags) -X github.com/swanchain/go-computing-provider/build.NetWorkTag=dev-proxima"
dev: computing-provider