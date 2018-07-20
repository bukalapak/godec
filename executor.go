package godec

import (
	"context"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

const (
	templatePath = "src/github.com/bukalapak/godec/template"
	outputPath   = "decorator"
)

type executor struct {
}

// NewExecutor creates an instance of executor.
func NewExecutor() Executor {
	return &executor{}
}

// Execute executes given godec interface to generate a new golang interface using given template.
// The generated file will be located in folder decorator, relative to current directory.
// The generated file's name will be `{tmpl.Name}/{intf.Name}.go`.
//
// Please, keep in mind that the template will be loaded from `$GOPATH/src/github.com/bukalapak/godec/template` folder.
// It will look for `{tmpl.Name}.go.tmpl` in that folder.
//
// For convenience, the generated file will automatically be formatted using `gofmt -s w`.
func (e *executor) Execute(ctx context.Context, intf *Interface, tmpl *Template) error {
	oPath := path.Join(outputPath, tmpl.Name)
	if err := os.MkdirAll(oPath, os.ModePerm); err != nil {
		return errors.Wrap(err, "couldn't make output directory")
	}

	fileName := path.Join(oPath, strings.ToLower(intf.Name)+".go")
	file, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, "couldn't create decorator file")
	}

	t, err := template.ParseFiles(path.Join(os.Getenv("GOPATH"), templatePath, tmpl.Name+".go.tmpl"))
	if err != nil {
		return errors.Wrap(err, "could't parse template file")
	}

	if err = t.Execute(file, intf); err != nil {
		return errors.Wrap(err, "couldn't make decorator from template")
	}

	return exec.CommandContext(ctx, "gofmt", "-s", "-w", fileName).Run()
}
