package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var Package, TypeName, Type string

func init() {
	flag.StringVar(&Package, "p", "swrap", "package name")
	flag.StringVar(&TypeName, "n", "SWrap", "type name")
	flag.StringVar(&Type, "t", "byte", "type name")
	flag.Parse()
}

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
		Package:  Package,
		TypeName: TypeName,
		Type:     Type,
		Upper:    strings.Title(Type),
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
