package godec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bukalapak/godec"
)

func TestNewExecutor(t *testing.T) {
	executor := godec.NewExecutor()
	assert.Implements(t, (*godec.Executor)(nil), executor)
}
