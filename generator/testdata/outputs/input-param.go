// Code generated by sham. DO NOT EDIT.

package outputs

import inputs "github.com/jmalloc/sham/generator/testdata/inputs"

// InputParam is a test implementation of the inputs.InputParam interface.
type InputParam struct {
	inputs.InputParam

	// AnonFunc is an implementation of the Anon() method.
	// If it is non-nil, it takes precedence over x.InputParam.Anon().
	AnonFunc func(int)

	// SingleFunc is an implementation of the Single() method.
	// If it is non-nil, it takes precedence over x.InputParam.Single().
	SingleFunc func(int)

	// MultipleFunc is an implementation of the Multiple() method.
	// If it is non-nil, it takes precedence over x.InputParam.Multiple().
	MultipleFunc func(int, float64)

	// MultipleNamesFunc is an implementation of the MultipleNames() method.
	// If it is non-nil, it takes precedence over x.InputParam.MultipleNames().
	MultipleNamesFunc func(int, int)

	// VariadicFunc is an implementation of the Variadic() method.
	// If it is non-nil, it takes precedence over x.InputParam.Variadic().
	VariadicFunc func(...int)
}

func (x *InputParam) Anon(i0 int) {
	if x.AnonFunc != nil {
		x.AnonFunc(i0)
	}

	if x.InputParam != nil {
		x.InputParam.Anon(i0)
	}
}

func (x *InputParam) Single(v int) {
	if x.SingleFunc != nil {
		x.SingleFunc(v)
	}

	if x.InputParam != nil {
		x.InputParam.Single(v)
	}
}

func (x *InputParam) Multiple(a int, b float64) {
	if x.MultipleFunc != nil {
		x.MultipleFunc(a, b)
	}

	if x.InputParam != nil {
		x.InputParam.Multiple(a, b)
	}
}

func (x *InputParam) MultipleNames(a, b int) {
	if x.MultipleNamesFunc != nil {
		x.MultipleNamesFunc(a, b)
	}

	if x.InputParam != nil {
		x.InputParam.MultipleNames(a, b)
	}
}

func (x *InputParam) Variadic(args ...int) {
	if x.VariadicFunc != nil {
		x.VariadicFunc(args...)
	}

	if x.InputParam != nil {
		x.InputParam.Variadic(args...)
	}
}
