package godec

import "context"

type parser struct {
}

func (p *parser) Parse(ctx context.Context, file File) (Interface, error) {
	return Interface{}, nil
}
