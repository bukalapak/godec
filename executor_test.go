package godec_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bukalapak/godec"
)

func TestNewExecutor(t *testing.T) {
	executor := godec.NewExecutor()
	assert.Implements(t, (*godec.Executor)(nil), executor)
}

func Test_executor_Execute_Success(t *testing.T) {
	defer os.RemoveAll("decorator")

	intf := &godec.Interface{
		Name:        "AnInterface",
		Package:     "something",
		PackagePath: "github.com/you/something",
	}

	template := &godec.Template{
		Name: "canceler",
	}

	executor := godec.NewExecutor()

	// template found
	err := executor.Execute(context.Background(), intf, template)
	assert.Nil(t, err)
}

func Test_executor_Execute_Error(t *testing.T) {
	defer os.RemoveAll("decorator")

	intf := &godec.Interface{
		Name:        "AnInterface",
		Package:     "something",
		PackagePath: "github.com/you/something",
	}

	executor := godec.NewExecutor()

	// template not found
	notFoundTemplate := &godec.Template{
		Name: "404Template",
	}
	err := executor.Execute(context.Background(), intf, notFoundTemplate)
	assert.NotNil(t, err)
}
