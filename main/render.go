package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	// read template file
	b, err := ioutil.ReadFile("./swrap_template")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)

	// define params
	var param = struct {
		Package, TypeName, Type, Upper string
	}{
		Package:  "swrap",
		TypeName: "SWrap",
		Type:     "byte",
		Upper:    "Byte",
	}

	// open output file
	fd, err := os.Create("./swrap.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	// execute template
	t := template.Must(template.New("swrap").Parse(str))
	err = t.Execute(fd, param)
	if err != nil {
		log.Fatal(err)
	}
}
