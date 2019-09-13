package main

import (
	"log"
	"os"
	"text/template"
)

type meal struct {
	Meal  string
	Itens []item
}

type item struct {
	Name  string
	Price float64
}

type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := []restaurant{
		restaurant{
			Name: "Restaurant 1",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Itens: []item{
						item{
							Name:  "Item 1",
							Price: 1.99,
						},
						item{
							Name:  "Item 2",
							Price: 2.99,
						},
						item{
							Name:  "Item 3",
							Price: 3.99,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Itens: []item{
						item{
							Name:  "Item 4",
							Price: 4.99,
						},
						item{
							Name:  "Item 5",
							Price: 5.99,
						},
						item{
							Name:  "Item 6",
							Price: 6.99,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Itens: []item{
						item{
							Name:  "Item 7",
							Price: 7.99,
						},
						item{
							Name:  "Item 8",
							Price: 8.99,
						},
						item{
							Name:  "Item 9",
							Price: 9.99,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
