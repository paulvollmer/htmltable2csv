VERSION=0.1.0

all: lint test

build:
	@go build

lint:
	@go fmt ./...
	@golint ./...

test: build
	@./htmltable2csv -v
	@./htmltable2csv -url "https://www.w3schools.com/html/html_tables.asp" -selector "#customers > tbody > tr" -csv data.csv

release:
	git tag -a v${VERSION} -m "Version ${VERSION}"
	git push origin v${VERSION}
	goreleaser

.PHONY: all lint build test release
