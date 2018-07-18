package godec

import (
	"context"

	"github.com/pkg/errors"
	goparser "github.com/zpatrick/go-parser"
)

type parser struct {
}

func (p *parser) Parse(ctx context.Context, file File) (Interface, error) {
	f, err := goparser.ParseSingleFile(file.Location)
	if err != nil {
		errors.Wrap(err, "couldn't parse file "+file.Location)
	}
}
