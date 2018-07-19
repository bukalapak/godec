package godec

import (
	"context"
	"fmt"
	"regexp"

	"github.com/bukalapak/godec"
	"github.com/pkg/errors"
	goparser "github.com/zpatrick/go-parser"
)

// Parser is a struct that implement godec.Parser interface.
type Parser struct {
}

// NewParser creates an instance of Parser.
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses godec file to godec interface.
func (p *Parser) Parse(ctx context.Context, file godec.File) (godec.Interface, error) {
	f, err := goparser.ParseSingleFile(file.Location)
	if err != nil {
		return godec.Interface{}, errors.Wrap(err, "couldn't parse file")
	}

	pkg, err := f.ImportPath()
	if err != nil {
		return godec.Interface{}, errors.Wrap(err, "couldn't get import path")
	}

	i, err := p.findInterface(f, file.Interface)
	if err != nil {
		return godec.Interface{}, err
	}

	intf := godec.Interface{
		Name:        i.Name,
		Package:     f.Package,
		PackagePath: pkg,
		Methods:     p.findMethods(f.Package, i),
	}

	return intf, nil
}

func (p *Parser) findInterface(f *goparser.GoFile, name string) (*goparser.GoInterface, error) {
	for _, intf := range f.Interfaces {
		if intf.Name == name {
			return intf, nil
		}
	}

	return nil, fmt.Errorf("interface %s not found", name)
}

func (p *Parser) findMethods(pkg string, intf *goparser.GoInterface) []godec.Method {
	var methods []godec.Method

	for _, m := range intf.Methods {
		method := godec.Method{Name: m.Name}

		for _, prm := range m.Params {
			param := godec.DataType{
				Name: "x",
				Type: p.getType(pkg, prm),
			}
			method.Params = append(method.Params, param)
		}

		for _, res := range m.Results {
			result := godec.DataType{
				Type:      p.getType(pkg, res),
				ZeroValue: p.getZeroValue(pkg, res),
			}
			method.ReturnValues = append(method.ReturnValues, result)
		}

		methods = append(methods, method)
	}

	return methods
}

func (p *Parser) getType(pkg string, t *goparser.GoType) string {
	if m, err := regexp.MatchString("[.]", t.Underlying); err != nil && m {
		return t.Type
	} else if t.Type[0] == '*' {
		return "*" + pkg + "." + t.Type[1:]
	} else {
		return pkg + "." + t.Type
	}
}

func (p *Parser) getZeroValue(pkg string, t *goparser.GoType) string {
	if t.Underlying[:6] == "struct" {
		return p.getType(pkg, t) + "{}"
	} else if m, err := regexp.MatchString(".*int.*", t.Underlying); err != nil && m {
		return "0"
	} else if m, err := regexp.MatchString(".*float.*", t.Underlying); err != nil && m {
		return "0"
	} else if t.Underlying == "string" {
		return "\"\""
	} else {
		return "nil"
	}
}
