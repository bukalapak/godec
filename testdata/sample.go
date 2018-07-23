package testdata

type Struct struct {
}

type Something struct {
}

type Sample interface {
	A(s Struct) (int, error)
	B(s Struct) (Struct, error)
	C(s Struct) float64
	D(s Struct) string
	E(s Struct, st *Something) *Something
}
