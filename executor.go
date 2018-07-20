package godec

import "context"

type executor struct {
}

// Execute executes given godec interface to create a new golang interface using given template.
// The new file will be located in folder decorator, relative to current directory.
// The new file's name will be `{template.Name}/{intf.Name}.go`.
//
// Please, keep in mind that the template will be loaded from `$GOPATH/src/github.com/bukalapak/godec/template` folder.
// It will look for `{template.Name}.go.tmpl` in that folder.
func (e *executor) Execute(ctx context.Context, intf *Interface, template *Template) error {

}
