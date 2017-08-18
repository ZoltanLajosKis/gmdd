BINARY    := gmdd
VERSION   := $(shell git describe --tags --abbrev=0)
REVISION  := $(shell git rev-parse --short HEAD)
LDFLAGS   := -ldflags "-X \"main.version=${VERSION}\" -X \"main.revision=${REVISION}\""

SOURCES   := $(shell find . -name '*.go' | grep -v './vendor/')
PACKAGES  := $(shell go list ./... | grep -v '/vendor/')

ASSETS_GEN := generate/assets.go
ASSETS_SRC := $(shell find generate/assets/ -type f)
ASSETS_OUT := assets/assets.go

TEMPLATES_GEN := generate/templates.go
TEMPLATES_SRC := $(shell find generate/templates/ -type f)
TEMPLATES_PAT := templates/%.go
TEMPLATES_OUT := $(patsubst generate/templates/%, ${TEMPLATES_PAT}, ${TEMPLATES_SRC})


${BINARY}: ${ASSETS_OUT} ${TEMPLATES_OUT} ${SOURCES} vendor
		go build ${LDFLAGS} -o ${BINARY}

${ASSETS_OUT}: ${ASSETS_GEN} ${ASSETS_SRC} vendor
		go run ${ASSETS_GEN}

${TEMPLATES_PAT}: ${TEMPLATES_GEN} ${TEMPLATES_SRC} vendor
		go run ${TEMPLATES_GEN}

vendor: | dep
		dep ensure -v

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
		go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: update-deps
update-deps: dep
		dep ensure -update -v
		@touch vendor


.PHONY: test
test:
		go test -cover -v $(PACKAGES)


.PHONY: clean
clean:
		rm -f ${BINARY}

.PHONY: clean-all
clean-all: clean
		rm -rf vendor/
		rm -rf assets/ templates/

