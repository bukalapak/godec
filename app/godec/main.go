package main

import (
	"context"
	"os"

	"github.com/bukalapak/godec"
)

func main() {
	file := &godec.File{
		Location:  os.Getenv("GOFILE"),
		Interface: os.Args[1],
	}

	var templates []*godec.Template
	for _, t := range os.Args[2:] {
		template := &godec.Template{Name: t}
		templates = append(templates, template)
	}

	parser := godec.NewParser()
	executor := godec.NewExecutor()

	decorator := godec.NewDecorator(parser, executor)
	decorator.Decorate(context.Background(), file, templates...)
}
