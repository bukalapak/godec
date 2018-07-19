package godec

import (
	"context"
	"fmt"

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

	intf, err := p.findInterface(f, file.Interface)
	if err != nil {
		return Interface{}, err
	}
}

func (p *parser) findInterface(f *goparser.GoFile, name string) (*goparser.GoInterface, error) {
	for _, intf := range f.Interfaces {
		if intf.Name == name {
			return intf, nil
		}
	}

	return nil, fmt.Errorf("interface %s not found", name)
}
