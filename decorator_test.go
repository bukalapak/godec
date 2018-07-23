package godec_test

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/bukalapak/godec"
	"github.com/stretchr/testify/assert"
)

func TestNewDecorator(t *testing.T) {
	parser := godec.NewParser()
	executor := godec.NewExecutor()
	decorator := godec.NewDecorator(parser, executor)
	assert.Implements(t, (*godec.Decorator)(nil), decorator)
}

func Test_decorator_Decorate_Success(t *testing.T) {
	file := &godec.File{
		Location:  path.Join(os.Getenv("GOPATH"), "src/github.com/bukalapak/godec/testdata/sample.go"),
		Interface: "Sample",
	}

	template := &godec.Template{
		Name: "canceler",
	}

	parser := godec.NewParser()
	executor := godec.NewExecutor()
	decorator := godec.NewDecorator(parser, executor)

	err := decorator.Decorate(context.Background(), file, template)
	assert.Nil(t, err)
}
