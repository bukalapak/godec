package testdata

import (
	"context"

	"github.com/bukalapak/godec"
)

type Struct struct {
}

type Something struct {
}

type Nanika interface {
	A(s Struct) (int, error)
}

//go:generate godec Sample instrumentation
type Sample interface {
	A(s Struct) (int, error)
	B(s Struct) (Struct, error)
	C(s Struct) float64
	D(s Struct) string
	E(s Struct, st *Something) *Something
	F(s Struct, f godec.File) *godec.Interface
	G(ctx context.Context, s string) *Struct
	H(ctx context.Context, i int) Nanika
}
