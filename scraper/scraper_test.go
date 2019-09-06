package htmltable2csv

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScraper(t *testing.T) {
	t.Run("source file", func(t *testing.T) {
		scraper := Scraper{}
		scraper.Source = "./fixture/test1.html"
		scraper.Selector = "table > tbody > tr"
		data, err := scraper.Scrape()
		if err != nil {
			t.Error(err)
		}
		dataEqual(t, data)
	})

	t.Run("source url", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte(`<table>
			<thead>
				<tr>
					<td>key</td>
					<td>value</td>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>foo</td>
					<td>1</td>
				</tr>
				<tr>
					<td>bar</td>
					<td>2</td>
				</tr>
				<tr>
					<td>baz</td>
					<td>3</td>
				</tr>
			</tbody>
		</table>`))
		}))
		defer server.Close()

		scraper := Scraper{}
		scraper.Source = server.URL
		scraper.Selector = "table > tbody > tr"
		data, err := scraper.Scrape()
		if err != nil {
			t.Error(err)
		}
		dataEqual(t, data)
	})
}

func dataEqual(t *testing.T, data [][]string) {
	if len(data) != 3 {
		t.Error("data not equal")
	}

	if len(data[0]) != 2 {
		t.Error("data[0] not equal")
	}
	if data[0][0] != "foo" {
		t.Error("data[0][0] not equal")
	}
	if data[0][1] != "1" {
		t.Error("data[0][1] not equal")
	}

	if len(data[1]) != 2 {
		t.Error("data[1] not equal")
	}
	if data[1][0] != "bar" {
		t.Error("data[1][0] not equal")
	}
	if data[1][1] != "2" {
		t.Error("data[1][1] not equal")
	}

	if len(data[2]) != 2 {
		t.Error("data[2] not equal")
	}
	if data[2][0] != "baz" {
		t.Error("data[2][0] not equal")
	}
	if data[2][1] != "3" {
		t.Error("data[2][0] not equal")
	}
}
