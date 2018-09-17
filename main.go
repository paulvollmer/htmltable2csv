package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/paulvollmer/htmltable2csv/scraper"
)

const version = "0.1.0"

func usage() {
	fmt.Println("Usage: htmltable2csv [flags]")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  htmltable2csv -url 'https://www.w3schools.com/html/html_tables.asp' -selector '#customers > tbody > tr' -csv data.csv")
	fmt.Println("")
}

func main() {
	flagVersion := flag.Bool("v", false, "Print the version and exit")
	flagURL := flag.String("url", "", "The website url")
	flagSelector := flag.String("selector", "", "The table css selector")
	flagCSV := flag.String("csv", "", "The csv filename. if empty, print csv to stdout")
	flag.Usage = usage
	flag.Parse()

	if *flagVersion {
		fmt.Printf("v%s\n", version)
		os.Exit(0)
	}

	if *flagURL == "" {
		fmt.Println("Flag -url cannot be empty")
		os.Exit(1)
	}

	if *flagSelector == "" {
		fmt.Println("Flag -selector cannot be empty")
		os.Exit(1)
	}

	var err error
	scraper := htmltable2csv.Scraper{}
	scraper.URL = *flagURL
	scraper.Selector = *flagSelector
	_, err = scraper.Scrape()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *flagCSV == "" {
		err = scraper.CSV(os.Stdout)
	} else {
		err = scraper.WriteCSV(*flagCSV)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
