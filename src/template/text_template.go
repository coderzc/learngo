package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseTemple(tmpl *template.Template, data interface{}) string {
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)
	CheckErr(err)
	return buf.String()
}

func main() {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	CheckErr(err)
	out := parseTemple(tmpl, sweaters)
	fmt.Println(out)
}
