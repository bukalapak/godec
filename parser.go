package godec

import (
	"context"
	"fmt"
	"regexp"

	"github.com/pkg/errors"
	goparser "github.com/zpatrick/go-parser"
)

type parser struct {
}

// NewParser creates an instance of Parser.
func NewParser() Parser {
	return &parser{}
}

// Parse parses godec file to godec interface.
func (p *parser) Parse(ctx context.Context, file *File) (*Interface, error) {
	f, err := goparser.ParseSingleFile(file.Location)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't parse file")
	}

	pkg, err := f.ImportPath()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get import path")
	}

	i, err := p.findInterface(f, file.Interface)
	if err != nil {
		return nil, err
	}

	intf := &Interface{
		Name:        i.Name,
		Package:     f.Package,
		PackagePath: pkg,
		Methods:     p.findMethods(f.Package, i),
	}

	return intf, nil
}

func (p *parser) findInterface(f *goparser.GoFile, name string) (*goparser.GoInterface, error) {
	for _, intf := range f.Interfaces {
		if intf.Name == name {
			return intf, nil
		}
	}

	return nil, fmt.Errorf("interface %s not found", name)
}

func (p *parser) findMethods(pkg string, intf *goparser.GoInterface) []Method {
	var methods []Method

	for _, m := range intf.Methods {
		method := Method{Name: m.Name}

		for idx, prm := range m.Params {
			param := DataType{
				Name: p.intToString(idx),
				Type: p.getType(pkg, prm),
			}
			method.Params = append(method.Params, param)
		}

		for _, res := range m.Results {
			result := DataType{
				Type:      p.getType(pkg, res),
				ZeroValue: p.getZeroValue(pkg, res),
			}
			method.ReturnValues = append(method.ReturnValues, result)
		}

		methods = append(methods, method)
	}

	return methods
}

func (p *parser) getType(pkg string, t *goparser.GoType) string {
	if found, err := regexp.MatchString(`^(\*|)`+pkg+`.`, t.Underlying); err == nil && found {
		return t.Underlying
	}
	return t.Type
}

func (p *parser) getZeroValue(pkg string, t *goparser.GoType) string {
	if len(t.Underlying) >= 6 && t.Underlying[:6] == "struct" {
		return p.getType(pkg, t) + "{}"
	} else if found, err := regexp.MatchString(".*int.*", t.Type); err == nil && found {
		return "0"
	} else if found, err := regexp.MatchString(".*float.*", t.Type); err == nil && found {
		return "0"
	} else if t.Type == "string" {
		return "\"\""
	} else {
		return "nil"
	}
}

func (p *parser) intToString(i int) string {
	return string('a' + i)
}
