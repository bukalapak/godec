package godec

import (
	"context"
	"os"
	"path"
	"text/template"
)

const (
	templatePath = "src/github.com/bukalapak/godec/template"
)

type executor struct {
}

// Execute executes given godec interface to create a new golang interface using given template.
// The new file will be located in folder decorator, relative to current directory.
// The new file's name will be `{tmpl.Name}/{intf.Name}.go`.
//
// Please, keep in mind that the template will be loaded from `$GOPATH/src/github.com/bukalapak/godec/template` folder.
// It will look for `{tmpl.Name}.go.tmpl` in that folder.
func (e *executor) Execute(ctx context.Context, intf *Interface, tmpl *Template) error {
	t, err := template.ParseFiles(path.Join(os.Getenv("GOPATH"), templatePath, tmpl.Name+".go.tmpl"))
	if err != nil {
		return err
	}
}
