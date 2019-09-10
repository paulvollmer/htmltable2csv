VERSION=0.2.1

all: lint test

build:
	@go build

lint:
	@go fmt ./...
	@golint ./...

test: test-src test-cli
test-src:
	@go test ./...
test-cli: build
	@./htmltable2csv -v
	@./htmltable2csv -source "./scraper/fixture/test1.html"                   -selector "table > tbody > tr"      -csv data_file.csv
	@./htmltable2csv -source "https://www.w3schools.com/html/html_tables.asp" -selector "#customers > tbody > tr" -csv data_url.csv

release:
	git tag -a v${VERSION} -m "Version ${VERSION}"
	git push origin v${VERSION}
	goreleaser --rm-dist
release-dry:
	goreleaser --skip-publish --skip-validate --snapshot

changelog:
	./node_modules/.bin/auto-changelog -p --template keepachangelog --breaking-pattern breaking && git add CHANGELOG.md

.PHONY: all lint build test test-all test-cli release release-dry changelog
