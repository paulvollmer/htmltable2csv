# htmltable2csv [![Build Status](https://travis-ci.org/paulvollmer/htmltable2csv.svg?branch=master)](https://travis-ci.org/paulvollmer/htmltable2csv) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/paulvollmer/htmltable2csv/blob/master/LICENSE)

## Introduction
`htmltable2csv` is a tool to parse a html table and store the data as csv. It can be written to a file or print out to `stdout`.

## Installation
##### Manually
Download your preferred flavor from the [releases page](https://github.com/paulvollmer/htmltable2csv/releases) and install manually.

##### Using go get
```
go get -u github.com/paulvollmer/htmltable2csv
```

## Usage
```
Usage: htmltable2csv [flags]

Flags:
  -csv string
    	The csv filename. if empty, print csv to stdout
  -selector string
    	The table css selector
  -url string
    	The website url
  -v	Print the version and exit

Examples:
  htmltable2csv -url 'https://www.w3schools.com/html/html_tables.asp' -selector '#customers > tbody > tr' -csv data.csv
```

## License
[MIT license](LICENSE)
