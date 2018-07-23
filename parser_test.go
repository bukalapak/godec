package godec_test

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/bukalapak/godec"
	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	parser := godec.NewParser()
	assert.Implements(t, (*godec.Parser)(nil), parser)
}

func Test_parser_Parse_Success(t *testing.T) {
	parser := godec.NewParser()

	file := &godec.File{
		Location:  path.Join(os.Getenv("GOPATH"), "src/github.com/bukalapak/godec/godec.go"),
		Interface: "Parser",
	}

	intf, err := parser.Parse(context.Background(), file)
	assert.Nil(t, err)

	assert.Equal(t, "Parser", intf.Name)
	assert.Equal(t, "godec", intf.Package)
	assert.Equal(t, "github.com/bukalapak/godec", intf.PackagePath)
	assert.Equal(t, "Parse", intf.Methods[0].Name)
	assert.Equal(t, "a", intf.Methods[0].Params[0].Name)
	assert.Equal(t, "context.Context", intf.Methods[0].Params[0].Type)
	assert.Equal(t, "b", intf.Methods[0].Params[1].Name)
	assert.Equal(t, "*godec.File", intf.Methods[0].Params[1].Type)
	assert.Equal(t, "*godec.Interface", intf.Methods[0].ReturnValues[0].Type)
	assert.Equal(t, "nil", intf.Methods[0].ReturnValues[0].ZeroValue)
	assert.Equal(t, "error", intf.Methods[0].ReturnValues[1].Type)
	assert.Equal(t, "nil", intf.Methods[0].ReturnValues[1].ZeroValue)
}
