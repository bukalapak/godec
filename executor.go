package godec

import (
	"context"
	"fmt"
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
// For convenience, the generated file will automatically be formatted using `goimports -w <file name>`.
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
	t, err := template.New(tmpl.Name + ".go.tmpl").
		Funcs(templateFuncMap()).
		ParseFiles(path.Join(os.Getenv("GOPATH"), templatePath, tmpl.Name+".go.tmpl"))
	if err != nil {
		return errors.Wrap(err, "couldn't parse template file")
	}

	t = t.Funcs(templateFuncMap())
	if err = t.Execute(file, intf); err != nil {
		return errors.Wrap(err, "couldn't make decorator from template")
	}

	if err := exec.CommandContext(ctx, "goimports", "-w", fileName).Run(); err != nil {
		return err
	}

	return exec.CommandContext(ctx, "gofmt", "-w", fileName).Run()
}

// templateFuncMap return a functions we need in template
func templateFuncMap() template.FuncMap {
	fun := template.FuncMap{
		// Search error in template variable
		"captureLatency": func(methodName string, values []DataType) string {
			for index, v := range values {
				if v.Type == "error" {
					return fmt.Sprintf("captureLatency( \"%s\" , start, type%d )", methodName, index)
				}
			}

			return fmt.Sprintf("captureLatency( \"%s\" , start, nil )", methodName)
		},
	}

	return fun
}
