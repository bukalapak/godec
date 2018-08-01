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
		Location:  path.Join(os.Getenv("GOPATH"), "src/github.com/bukalapak/godec/testdata/sample.go"),
		Interface: "Sample",
	}

	intf, err := parser.Parse(context.Background(), file)
	assert.Nil(t, err)

	assert.Equal(t, "Sample", intf.Name)
	assert.Equal(t, "testdata", intf.Package)
	assert.Equal(t, "github.com/bukalapak/godec/testdata", intf.PackagePath)
	assert.Equal(t, "A", intf.Methods[0].Name)
	assert.Equal(t, "a", intf.Methods[0].Params[0].Name)
	assert.Equal(t, "github.com/bukalapak/godec/testdata.Struct", intf.Methods[0].Params[0].Type)
	assert.Equal(t, "int", intf.Methods[0].ReturnValues[0].Type)
	assert.Equal(t, "0", intf.Methods[0].ReturnValues[0].ZeroValue)
	assert.Equal(t, "error", intf.Methods[0].ReturnValues[1].Type)
	assert.Equal(t, "nil", intf.Methods[0].ReturnValues[1].ZeroValue)
}

func Test_parser_Parse_Error(t *testing.T) {
	parser := godec.NewParser()

	// no file
	notFoundFile := &godec.File{
		Location: path.Join(os.Getenv("GOPATH"), "src/github.com/bukalapak/godec/testdata/something.go"),
	}

	intf, err := parser.Parse(context.Background(), notFoundFile)
	assert.NotNil(t, err)
	assert.Nil(t, intf)

	// no desired interface
	noDesiredInterface := &godec.File{
		Location:  path.Join(os.Getenv("GOPATH"), "src/github.com/bukalapak/godec/testdata/sample.go"),
		Interface: "someinterface",
	}

	intf, err = parser.Parse(context.Background(), noDesiredInterface)
	assert.NotNil(t, err)
	assert.Nil(t, intf)
}
