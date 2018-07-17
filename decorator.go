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

	// Methods are list of interface's methods.
	Methods []Method
}

// Method is a struct that represents golang method data.
type Method struct {
	// Name is method's name.
	Name string

	// Params are method's parameters.
	// It is assumed that the first parameter will always be a context.Context
	// and it shouldn't be provided in Params.
	// Other parameters except for context.Context should all be provided in Params accordingly.
	// They will be the second parameters and so on.
	Params []DataType

	// ReturnValues are method's return values.
	// It is assumed that the last return value will always be an error
	// and it shouldn't be provided in ReturnValues.
	// Other return values except for error should all be provided in ReturnValues accordingly.
	// They will be the first, second, until the N-1 return value.
	// Then, the last return value will always be an error.
	ReturnValues []DataType
}
