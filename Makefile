VERSION=`git describe --tags`
BUILDDATE=`date +%FT%T%z`
BUILDMODE=-buildmode=plugin
LDFLAGS=-ldflags="-X 'main.PluginVersion=${VERSION}' -X 'main.PluginBuilddate=${BUILDDATE}'"

PLATFORMS := linux/amd64 linux/arm64 linux/386 linux/arm darwin/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

.PHONY: help
help:  ## 🤔 Show help messages
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

update: ## 🔄 Update dependencies
	go mod tidy
	go get -u

build: ## 🚧 Build for local arch
	mkdir -p ./lib
	go build ${BUILDMODE} -o ./lib/coffee.so ${LDFLAGS} ./*.go

clean: ## 🧹 Remove previously build binaries
	rm -rf ./lib

pre-release:
	mkdir -p ./lib/release

release: pre-release $(PLATFORMS) ## 📦 Build for GitHub release
$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build ${BUILDMODE} -o ./lib/release/coffee-$(os)-$(arch).so ${LDFLAGS} ./*.go
