package main

import (
	"os"
	"text/template"
)

func Multiply(a, b float64) float64 {
	return a * b
}

func main() {
	tmpl := template.New("Foo")
	tmpl.Funcs(template.FuncMap{"multiply": Multiply})

	tmpl, err = tmpl.Parse(
		"Price: ${{ multiply .Price .Quantity | printf \"%.2f\"}}\n",
	)
	if err != nil {
		panic(err)
	}

	type Product struct {
		Price    float64
		Quantity float64
	}
	err = tmpl.Execute(os.Stdout, Product{
		Price:    12.3,
		Quantity: 2,
	})
	if err != nil {
		panic(err)
	}
}
