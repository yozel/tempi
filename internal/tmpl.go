package internal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type tempifile struct {
	Template string
	Values   map[string]interface{}
}

type templateContext struct {
	Values interface{}
}

func Render(tempiFilePath string, outputFilePaths []string, format string) {
	yamlFile, err := ioutil.ReadFile(tempiFilePath)
	if err != nil {
		log.Fatal(err)
	}

	c := tempifile{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	templateFile, err := ioutil.ReadFile(path.Join(path.Dir(tempiFilePath), c.Template))
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("tempi").Option("missingkey=error").Parse(string(templateFile))
	if err != nil {
		log.Fatal(err)
	}

	var output bytes.Buffer

	err = tmpl.Execute(&output, templateContext{Values: c.Values})
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, o := range outputFilePaths {
		t := strings.SplitN(o, ":", 2)
		format := t[0]
		path := t[1]

		var of io.Writer
		if path == "-" {
			of = os.Stdout
		} else {
			of, err = os.Create(path)
			if err != nil {
				log.Fatal(err)
			}
		}

		if format == "txt" {
			err = writeText(of, output.String())
		} else if format == "pdf" {
			err = writePdf(of, output.String())
		} else {
			err = errors.New(fmt.Sprintf("Format is not supported: %s", format))
		}

		if err != nil {
			log.Fatal(err)
		}

	}
}
