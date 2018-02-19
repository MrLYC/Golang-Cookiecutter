VERSION := 0.0.1

ROOTDIR := $(shell pwd)
APPNAME := {{cookiecutter.app_name}}
GODIR := /tmp/gopath
SRCDIR := ${GODIR}/src/${APPNAME}
TARGET := bin/${APPNAME}

GOENV := GOPATH=${GODIR}:${GOPATH} GO15VENDOREXPERIMENT=1

GO := ${GOENV} go
DEP := ${GOENV} dep

LDFLAGS := -X ${APPNAME}/config.Version=${VERSION}
DEBUGLDFLAGS := ${LDFLAGS} -X ${APPNAME}/config.Mode=debug
RELEASELDFLAGS := -w ${LDFLAGS} -X ${APPNAME}/config.Mode=release

.PHONY: release
release: ${SRCDIR}
	${GO} build -i -ldflags="${RELEASELDFLAGS}" -o ${TARGET} ${APPNAME}

.PHONY: build
build: ${SRCDIR}
	${GO} build -i -ldflags="${DEBUGLDFLAGS}" -o ${TARGET} ${APPNAME}

${SRCDIR}:
	mkdir -p bin
	mkdir -p `dirname "${SRCDIR}"`
	ln -s `dirname "${ROOTDIR}"`/${APPNAME} ${SRCDIR}

.PHONY: init
init: ${SRCDIR} update

.PHONY: update
update: ${SRCDIR}
	cd ${SRCDIR} && ${DEP} ensure || true
	find vendor -name 'testdata' -type d -exec rm -rf {} \; || true
	find vendor -name '*_test.go' -delete || true
	find vendor -type f \( ! -name '*.go' ! -name 'LICENSE' ! -name '*.s' ! -name '*.c' ! -name '*.cpp' \) -delete || true

.PHONY: test
test: ${SRCDIR}
	$(eval package ?= $(patsubst ./%,${APPNAME}/%,$(shell find "." -name "*_test.go" -not -path "./vendor/*" -not -path "./.*" -exec dirname {} \; | uniq)))
	${GOENV} go test ${package}

.PHONY: lint
lint:
	${GOENV} find . -type f -name "*.go" -not -path "./vendor/*" -exec golint {} \;

.PHONY: go-env
go-env:
	@go env
