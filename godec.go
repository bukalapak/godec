package godec

import (
	"context"
)

// Interface is a struct that represents golang interface data.
type Interface struct {
	// Name is interface's name.
	Name string

	// Package is interface's package name.
	Package string

	// PackagePath is full golang path to interface's package.
	// It is used for import purpose.
	PackagePath string

	// Methods are list of interface's methods.
	Methods []Method
}

// Method is a struct that represents golang method data.
type Method struct {
	// Name is method's name.
	Name string

	// Params are method's parameters.
	// It is assumed that the first parameter will always be a context.Context
	// and it shouldn't be provided in Params.
	// Other parameters except for context.Context should all be provided in Params accordingly.
	// They will be the second parameters and so on.
	Params []DataType

	// ReturnValues are method's return values.
	// It is assumed that the last return value will always be an error
	// and it shouldn't be provided in ReturnValues.
	// Other return values except for error should all be provided in ReturnValues accordingly.
	// They will be the first, second, until the N-1 return value.
	// Then, the last return value will always be an error.
	ReturnValues []DataType
}

// DataType is a struct that represent golang data type.
type DataType struct {
	// Name is variable name for data type.
	Name string

	// Type is type name such as string, error, map[string][string], interface{}, etc.
	Type string

	// ZeroValue is default value for data type.
	// It is used as default return value when an error is occurred.
	//
	// Example: zero value for int is 0, zero value for pointer is nil, etc.
	ZeroValue string
}

// File is a struct that represent a golang file that contains some interfaces.
type File struct {
	// Location is file's location in the system.
	// Example: github.com/bukalapak/godec/example.go.
	Location string

	// Interface is an interface that needs to be decorated.
	Interface string
}

// Template is a struct that holds decorator template data.
type Template struct {
	// Name is template's name.
	// Godec will use godec path as default path.
	// Thus, any template should be located in godec/template folder.
	//
	// Let say there is a template named canceler.go.tmpl in godec/template,
	// users should only use `canceler` as template's name.
	Name string
}

// Decorator is used to decorate a file.
type Decorator interface {
	// Decorate decorates the given file.
	Decorate(ctx context.Context, file *File, templates ...*Template) error
}

// Parser is used to parse file and find the desired interface.
type Parser interface {
	// Parse parse the given file.
	// It will find the desired interface and return it.
	Parse(ctx context.Context, file *File) (*Interface, error)
}

// Executor is used to execute interface using a template.
type Executor interface {
	// Execute executes given godec interface to create golang interface using given template.
	Execute(ctx context.Context, intf *Interface, template *Template) error
}
