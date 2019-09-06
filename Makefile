VERSION=0.2.0

all: lint test

build:
	@go build

lint:
	@go fmt ./...
	@golint ./...

test: build
	@./htmltable2csv -v
	@./htmltable2csv -source "./scraper/fixture/test1.html"                   -selector "table > tbody > tr"      -csv data_file.csv
	@./htmltable2csv -source "https://www.w3schools.com/html/html_tables.asp" -selector "#customers > tbody > tr" -csv data_url.csv

test-all:
	@go test ./...
	@make test

release:
	git tag -a v${VERSION} -m "Version ${VERSION}"
	git push origin v${VERSION}
	goreleaser --rm-dist
release-dry:
	goreleaser --skip-publish --skip-validate --snapshot

.PHONY: all lint build test test-all release release-dry
