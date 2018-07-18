package godec

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type decorator struct {
	parser Parser
}

func (d *decorator) Decorate(ctx context.Context, file File, templates ...Template) error {
	intf, err := d.parser.Parse(ctx, file)
	if err != nil {
		return errors.Wrap(err, "couldn't parse given file")
	}

	for _, template := range templates {
		err = d.decorate(intf, template)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed to create decorator for %s using template %s", file.Location, template.Location))
		}
	}

	return nil
}

func (d *decorator) decorate(intf Interface, template Template) error {
	return nil
}
