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

func Test_decorator_Decorate_Instrumentation_Success(t *testing.T) {
	file := &godec.File{
		Location:  path.Join(os.Getenv("GOPATH"), "src/github.com/bukalapak/godec/testdata/sample.go"),
		Interface: "Sample",
	}

	template := &godec.Template{
		Name: "instrumentation",
	}

	parser := godec.NewParser()
	executor := godec.NewExecutor()
	decorator := godec.NewDecorator(parser, executor)

	err := decorator.Decorate(context.Background(), file, template)
	assert.Nil(t, err)
}

func Test_decorator_Decorate_Error(t *testing.T) {
	// not found file
	notFoundfile := &godec.File{
		Location:  path.Join(os.Getenv("GOPATH"), "src/github.com/bukalapak/godec/testdata/notfound.go"),
		Interface: "Sample",
	}

	template := &godec.Template{
		Name: "canceler",
	}

	parser := godec.NewParser()
	executor := godec.NewExecutor()
	decorator := godec.NewDecorator(parser, executor)

	err := decorator.Decorate(context.Background(), notFoundfile, template)
	assert.NotNil(t, err)

	// not found template
	file := &godec.File{
		Location:  path.Join(os.Getenv("GOPATH"), "src/github.com/bukalapak/godec/testdata/sample.go"),
		Interface: "Sample",
	}

	notFoundTemplate := &godec.Template{
		Name: "notfoundtemplate",
	}

	err = decorator.Decorate(context.Background(), file, notFoundTemplate)
	assert.NotNil(t, err)
}
