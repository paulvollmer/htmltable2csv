<p align="center">
  <img src="https://fonts.gstatic.com/s/i/materialicons/business/v1/24px.svg" height="96"/>
  <h3 align="center">
    htmltable2csv
  </h3>
  <p align="center">
    <code>htmltable2csv</code> is a tool to parse a html table and store the data as csv, written to a <code>file</code> or print to <code>stdout</code>.
  </p>
  <p align="center">
    <a href="https://travis-ci.org/paulvollmer/htmltable2csv"><img alt="Travis" src="https://img.shields.io/travis/paulvollmer/htmltable2csv/master.svg?style=flat-square"> </a>
    <a href="https://github.com/paulvollmer/htmltable2csv/releases"><img alt="Software Release" src="https://img.shields.io/github/v/release/paulvollmer/htmltable2csv.svg?style=flat-square"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square"></a>
  </p>
</p>

---

## Installation
##### Using go get
```
go get -u github.com/paulvollmer/htmltable2csv
```

##### Manually
Download your preferred flavor from the [releases page](https://github.com/paulvollmer/htmltable2csv/releases) and install manually.

## Usage
```
Usage: htmltable2csv [flags]

Flags:
  -csv string
    	The csv filename. if empty, print csv to stdout
  -selector string
    	The table css selector
  -source string
    	The filepath or website url
  -trim
    	Trim the whitespace for each table column
  -v	Print the version and exit

Examples:
  htmltable2csv -url 'https://www.w3schools.com/html/html_tables.asp' -selector '#customers > tbody > tr' -csv data.csv
```

## License
[MIT License](LICENSE)
