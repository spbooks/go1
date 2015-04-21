package main

import (
	"fmt"
	"html/template"
	"os"
)

type Article struct {
	Name       string
	AuthorName string
	Draft      bool
}

func (a Article) Byline() string {
	return fmt.Sprintf("Written by %s", a.AuthorName)
}

func main() {
	//Example 1
	tmpl, err := template.New("Foo").Parse("<h1>Hello {{.}}</h1>\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, "World")
	if err != nil {
		panic(err)
	}

	// Example 2
	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
	}
	tmpl, err = template.New("Foo").Parse("'{{.Name}}' by {{.AuthorName}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}

	// Example 3
	article := map[string]string{
		"Name":       "The Go html/template package",
		"AuthorName": "Mal Curtis",
	}
	tmpl, err = template.New("Foo").Parse("'{{.Name}}' by {{.AuthorName}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, article)
	if err != nil {
		panic(err)
	}

	// Example 4
	tmpl, err = template.New("Foo").Parse("{{.Byline}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}

	// Example 5
	goArticle.Draft = true
	tmpl, err = template.New("Foo").Parse("{{.Name}}{{if .Draft}} (Draft){{end}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}

	// Example 6
	tmpl, err = template.New("Foo").Parse(`
	{{range .}}
		<p>{{.Name}} by {{.AuthorName}}</p>
	{{else}}
		<p>No published articles yet</p>
	{{end}}
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, []Article{})
	if err != nil {
		panic(err)
	}

	// Example 7
	tmpl, err = template.New("Foo").Parse(`
	{{define "ArticleResource"}}
		<p>{{.Name}} by {{.AuthorName}}</p>
	{{end}}

	{{define "ArticleLoop"}}
		{{range .}}
			{{template "ArticleResource" .}}
		{{else}}
			<p>No published articles yet</p>
		{{end}}
	{{end}}

	{{template "ArticleLoop" .}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, []Article{goArticle})
	if err != nil {
		panic(err)
	}

	// Example 8
	tmpl, err = template.New("Foo").Parse("Price: ${{printf \"%.2f\" .}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, 12.3)
	if err != nil {
		panic(err)
	}

	// Example 9
	tmpl, err = template.New("Foo").Parse("Price: ${{. | printf \"%.2f\"}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, 12.3)
	if err != nil {
		panic(err)
	}

	// Example 10
	type Product struct {
		Price    float64
		Quantity float64
	}
	tmpl = template.New("Foo")
	tmpl.Funcs(template.FuncMap{
		"multiply": Multiply,
	})

	tmpl, err = tmpl.Parse("Price: ${{ multiply .Price .Quantity | printf \"%.2f\"}}\n")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, Product{
		Price:    12.3,
		Quantity: 2,
	})
	if err != nil {
		panic(err)
	}

	// Example 11
	tmpl = template.New("Foo")
	tmpl.Funcs(template.FuncMap{
		"multiply": Multiply,
	})

	tmpl, err = tmpl.Parse(`
	{{$total := multiply .Price .Quantity}}
	Price: ${{ printf "%.2f" $total}}
	`)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, Product{
		Price:    12.3,
		Quantity: 2,
	})
	if err != nil {
		panic(err)
	}
}
func Multiply(a, b float64) float64 {
	return a * b
}
