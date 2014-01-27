package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {

	b, err := ioutil.ReadFile("./swrap_template")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)

	type Defun struct {
		Package, TypeName, Type, Upper string
	}
	var defun = Defun{
		Package:  "swrap",
		TypeName: "SWrap",
		Type:     "byte",
		Upper:    "Byte",
	}

	t := template.Must(template.New("swrap").Parse(str))

	fd, err := os.Create("./swrap.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	err = t.Execute(fd, defun)
	if err != nil {
		log.Fatal(err)
	}
}
