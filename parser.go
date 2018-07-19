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

func (p *parser) Parse(ctx context.Context, file File) (Interface, error) {
	f, err := goparser.ParseSingleFile(file.Location)
	if err != nil {
		return Interface{}, errors.Wrap(err, "couldn't parse file")
	}

	pkg, err := f.ImportPath()
	if err != nil {
		return Interface{}, errors.Wrap(err, "couldn't get import path")
	}

	i, err := p.findInterface(f, file.Interface)
	if err != nil {
		return Interface{}, err
	}

	intf := Interface{
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

		for _, prm := range m.Params {
			param := DataType{
				Name: "x",
				Type: p.getType(pkg, prm),
			}
			method.Params = append(method.Params, param)
		}
	}
}

func (p *parser) getType(pkg string, t *goparser.GoType) string {
	if m, err := regexp.MatchString("[.]", t.Underlying); err != nil && m {
		return t.Type
	} else if t.Type[0] == '*' {
		return "*" + pkg + "." + t.Type[1:]
	} else {
		return pkg + "." + t.Type
	}
}
