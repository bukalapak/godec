package godec_test

import (
	"testing"

	"github.com/bukalapak/godec"
	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	parser := godec.NewParser()
	assert.Implements(t, (*godec.Parser)(nil), parser)
}

func Test_parser_Parse(t *testing.T) {
}
