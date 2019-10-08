package htmltable2csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Scraper store the Source, Selector and collected Data
type Scraper struct {
	Source   string
	Selector string
	Data     [][]string
	Trim     bool
	Start    int
	Stop     int
}

// Scrape download and parse the table data
func (s *Scraper) Scrape() ([][]string, error) {
	var data = make([][]string, 0)

	var doc goquery.Document

	_, err := url.ParseRequestURI(s.Source)
	if err != nil {
		f, err := os.Open(s.Source)
		if err != nil {
			return data, err
		}
		defer f.Close()
		tmp, err := goquery.NewDocumentFromReader(f)
		if err != nil {
			return data, err
		}
		doc = *tmp
	} else {
		res, err := http.Get(s.Source)
		if err != nil {
			return data, err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return data, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		}
		tmp, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return data, err
		}
		doc = *tmp
	}

	// Find the table
	length := doc.Find(s.Selector).Length()
	if s.Start > length {
		return data, fmt.Errorf("cannot start looking at row %d when the table only has %d row(s)", s.Start, length)
	}
	if s.Stop == -1 {
		s.Stop = length
	}
	doc.Find(s.Selector).Each(func(i int, table *goquery.Selection) {
		index := table.Index()
		if index >= s.Start && index <= s.Stop {
			dataRow := make([]string, 0)
			table.Find("td").Each(func(j int, td *goquery.Selection) {
				text := td.Text()
				if s.Trim {
					text = strings.TrimSpace(text)
				}
				dataRow = append(dataRow, text)
			})
			data = append(data, dataRow)
		}
	})
	s.Data = data
	return data, nil
}

// CSV write the Data to the given io Writer
func (s *Scraper) CSV(w io.Writer) error {
	writer := csv.NewWriter(w)
	writer.WriteAll(s.Data)
	return writer.Error()
}

// WriteCSV write the Data to the given filename
func (s *Scraper) WriteCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return s.CSV(file)
}
