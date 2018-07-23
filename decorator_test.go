package godec_test

import (
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

func Test_decorator_Decorate(t *testing.T) {

}
