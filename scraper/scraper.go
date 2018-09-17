package htmltable2csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// Scraper store the URL, Selector and collected Data
type Scraper struct {
	URL      string
	Selector string
	Data     [][]string
}

// Scrape download and parse the table data
func (s *Scraper) Scrape() ([][]string, error) {
	var data = make([][]string, 0)
	res, err := http.Get(s.URL)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return data, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return data, err
	}
	// Find the table
	doc.Find(s.Selector).Each(func(i int, table *goquery.Selection) {
		dataRow := make([]string, 0)
		table.Find("td").Each(func(j int, td *goquery.Selection) {
			dataRow = append(dataRow, td.Text())
		})
		data = append(data, dataRow)
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
