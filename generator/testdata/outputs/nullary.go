// Code generated by sham. DO NOT EDIT.
// https://github.com/jmalloc/sham

package outputs

import inputs "github.com/jmalloc/sham/generator/testdata/inputs"

// Nullary is a test implementation of the inputs.Nullary interface.
type Nullary struct {
	// Nullary is the default implementation of the interface.
	// If it is nil, each method will return an error (or panic).
	inputs.Nullary

	// MethodFunc is an implementation of the Method() method.
	// If it is non-nil, it takes precedence over the embedded Nullary interface.
	MethodFunc func()
}

func (stub *Nullary) Method() {
	if stub.MethodFunc != nil {
		stub.MethodFunc()
	}

	if stub.Nullary != nil {
		stub.Nullary.Method()
	}
}
