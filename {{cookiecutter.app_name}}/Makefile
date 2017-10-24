VERSION := 0.0.1
ROOTDIR := $(shell pwd)
APPNAME := {{cookiecutter.app_name}}
SRCDIR := src/{{cookiecutter.app_name}}
TARGET := bin/{{cookiecutter.app_name}}

GOENV := GOPATH=${ROOTDIR}:${GOPATH} GO15VENDOREXPERIMENT=1

GO := ${GOENV} go
GLIDE := ${GOENV} glide

GLIDEYAML := ${ROOTDIR}/glide.yaml
GLIDELOCK := ${ROOTDIR}/glide.lock

LDFLAGS := -X ${APPNAME}/config.Version=${VERSION}
DEBUGLDFLAGS := ${LDFLAGS} -X ${APPNAME}/config.Mode=debug
RELEASELDFLAGS := -w ${LDFLAGS} -X ${APPNAME}/config.Mode=release

.PHONY: release
release: ${SRCDIR} ${GLIDELOCK}
	${GO} build -i -ldflags="${RELEASELDFLAGS}" -o ${TARGET} {{cookiecutter.app_name}}

.PHONY: build
build: ${SRCDIR} ${GLIDELOCK}
	${GO} build -i -ldflags="${DEBUGLDFLAGS}" -o ${TARGET} {{cookiecutter.app_name}}

${SRCDIR}:
	mkdir -p bin
	mkdir -p src
	ln -s `dirname "${ROOTDIR}"`/${APPNAME} src/

${GLIDEYAML}:
	${GLIDE} init

${GLIDELOCK}: ${SRCDIR} ${GLIDEYAML}
	${GLIDE} install
	touch ${GLIDELOCK}

.PHONY: init
init: ${SRCDIR} ${GLIDEYAML}

.PHONY: update
update: ${SRCDIR}
	${GLIDE} update
	find vendor -name 'testdata' -type d -exec rm -rf {} \; || true
	find vendor -name '*_test.go' -delete || true
	find vendor -type f \( ! -name '*.go' ! -name 'LICENSE' ! -name '*.s' ! -name '*.c' ! -name '*.cpp' \) -delete || true


.PHONY: install
install: ${GLIDELOCK}

.PHONY: test
test: init
	$(eval package ?= $(patsubst ./%,${APPNAME}/%,$(shell find "." -name "*_test.go" -not -path "./vendor/*" -not -path "./src/*" -not -path "./.*" -exec dirname {} \; | uniq)))
	${GOENV} go test ${package}

.PHONY: lint
lint:
	${GOENV} find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./src/*" -exec golint {} \;
