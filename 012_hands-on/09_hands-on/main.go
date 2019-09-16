package main

import (
	"strconv"
	"encoding/csv"
	"os"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Data struct {
	Date time.Time
	Open float64
}

var dataTemplate *template.Template

func init() {
	dataTemplate = template.Must(template.ParseFiles("data.gohtml"))
}

func main() {
	http.HandleFunc("/", serveData)
	http.ListenAndServe(":8080", nil)
}

func serveData(res http.ResponseWriter, req *http.Request) {

	// parse csv
	records := parseData("table.csv")

	// execute template
	err := dataTemplate.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}

}

func parseData(fileName string) []Data {
	src, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	csvReader := csv.NewReader(src)
	dataRows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Data, 0, len(dataRows))

	for i, row := range dataRows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Data{
			Date: date,
			Open: open,
		})
	}

	return records

}
