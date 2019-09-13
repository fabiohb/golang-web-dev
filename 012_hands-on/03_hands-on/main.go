package main

import (
	"log"
	"os"
	"text/template"
)

type region struct {
	Name   string
	Hotels []hotel
}

type hotel struct {
	Name, Address, City, Zip string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	regions := []region{
		region{
			Name: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Southern Hotel 1",
					Address: "Addr 1",
					City:    "City 1",
					Zip:     "000000000",
				},
				hotel{
					Name:    "Southern Hotel 2",
					Address: "Addr 2",
					City:    "City 2",
					Zip:     "000000000",
				},
			},
		},
		region{
			Name: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Central Hotel 1",
					Address: "Addr 1",
					City:    "City 1",
					Zip:     "000000000",
				},
				hotel{
					Name:    "Central Hotel 2",
					Address: "Addr 2",
					City:    "City 2",
					Zip:     "000000000",
				},
			},
		},
		region{
			Name: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "Northern Hotel 1",
					Address: "Addr 1",
					City:    "City 1",
					Zip:     "000000000",
				},
				hotel{
					Name:    "Northern Hotel 2",
					Address: "Addr 2",
					City:    "City 2",
					Zip:     "000000000",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, regions)
	if err != nil {
		log.Fatalln(err)
	}
}
