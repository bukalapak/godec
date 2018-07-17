package godec

// Interface is a struct that represents golang interface data.
type Interface struct {
	// Name is interface's name.
	Name string

	// Package is interface's package name.
	Package string

	// PackagePath is full golang path to interface's package.
	// It is used for import purpose.
	PackagePath string

	// Methods is list of interface's methods.
	Methods []Method
}
