VERSION := 0.0.1
ROOTDIR := $(shell pwd)
APPDIR := {{cookiecutter.app_name}}
SRCDIR := src/{{cookiecutter.app_name}}
TARGET := bin/{{cookiecutter.app_name}}

GOENV := GOPATH=${ROOTDIR}:${GOPATH} GO15VENDOREXPERIMENT=1

GO := ${GOENV} go
GLIDE := ${GOENV} glide

GLIDEYAML := ${ROOTDIR}/glide.yaml
GLIDELOCK := ${ROOTDIR}/glide.lock

LDFLAGS := -X ${APPDIR}/config.Version=${VERSION}
DEBUGLDFLAGS := ${LDFLAGS} -X ${APPDIR}/config.Mode=debug
RELEASELDFLAGS := -w ${LDFLAGS} -X ${APPDIR}/config.Mode=release

.PHONY: release
release: ${SRCDIR} ${GLIDELOCK}
	${GO} build -i -ldflags="${RELEASELDFLAGS}" -o ${TARGET} {{cookiecutter.app_name}}

.PHONY: build
build: ${SRCDIR} ${GLIDELOCK}
	${GO} build -i -ldflags="${DEBUGLDFLAGS}" -o ${TARGET} {{cookiecutter.app_name}}

${SRCDIR}:
	mkdir -p bin
	mkdir -p src
	ln -s ${ROOTDIR} src/

${GLIDEYAML}:
	${GLIDE} init

${GLIDELOCK}: ${SRCDIR} ${GLIDEYAML}
	${GLIDE} install

.PHONY: init
init: ${SRCDIR} ${GLIDEYAML}

.PHONY: update
update: ${SRCDIR}
	${GLIDE} update

.PHONY: install
install: ${GLIDELOCK}
