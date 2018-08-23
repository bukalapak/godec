package testdata

import (
	"context"

	"github.com/bukalapak/godec"
)

type ABC struct {
}

type Something struct {
}

type Nanika interface {
	A(s ABC) (int, error)
}

//go:generate godec Sample instrumentation
type Sample interface {
	A(s ABC) (int, error)
	B(s ABC) (ABC, error)
	C(s ABC) float64
	D(s ABC) string
	E(s ABC, st *Something) *Something
	F(s ABC, f godec.File) *godec.Interface
	G(ctx context.Context, s string) *ABC
	H(ctx context.Context, i int) Nanika
}
