package godec

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type decorator struct {
	parser   Parser
	executor Executor
}

// NewDecorator returns an instance of decorator.
func NewDecorator(parser Parser, executor Executor) Decorator {
	return &decorator{
		parser:   parser,
		executor: executor,
	}
}

func (d *decorator) Decorate(ctx context.Context, file *File, templates ...*Template) error {
	intf, err := d.parser.Parse(ctx, file)
	if err != nil {
		return errors.Wrap(err, "couldn't parse given file")
	}

	for _, template := range templates {
		err = d.executor.Execute(ctx, intf, template)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to create decorator for %s using template %s", file.Location, template.Name))
		}
	}

	return nil
}
