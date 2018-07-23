package godec_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bukalapak/godec"
)

func TestNewExecutor(t *testing.T) {
	executor := godec.NewExecutor()
	assert.Implements(t, (*godec.Executor)(nil), executor)
}

func Test_executor_Execute(t *testing.T) {
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

	// template not found
	notFoundTemplate := &godec.Template{
		Name: "404Template",
	}
	err = executor.Execute(context.Background(), intf, notFoundTemplate)
	assert.NotNil(t, err)
}
