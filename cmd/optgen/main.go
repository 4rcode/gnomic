package main

import (
	"bytes"
	_ "embed"
	"flag"
	"go/format"
	"log"
	"os"
	"text/template"
)

var fatal = log.Fatal

//go:embed optgen.tmpl
var templateText string

func main() {
	err := Run(os.Args[1:]...)

	if err != nil {
		fatal(err)
	}
}

func Run(args ...string) error {
	var flags _flags

	err := flags.parse(args...)

	if err != nil {
		if err == flag.ErrHelp {
			return nil
		}

		return err
	}

	var buffer bytes.Buffer

	err = template.Must(
		new(template.Template).
			Parse(templateText),
	).Execute(&buffer, flags)

	if err != nil {
		return err
	}

	bytes := buffer.Bytes()
	bytes, err = format.Source(bytes)

	if err != nil {
		return err
	}

	return os.WriteFile(flags.fileName, bytes, 0666)
}
