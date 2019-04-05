package testing

import "testing"

// Asserted makes testing.T asserted.
func Asserted(t *testing.T) *asserted {
	return (*asserted)(t)
}

type asserted testing.T

func (t *asserted) True() bool {
	return false
}

func (t *asserted) False() bool {
	return false
}

func (t *asserted) Error() bool {
	return false
}

func (t *asserted) EqualError() bool {
	return false
}

func (t *asserted) MatchError() bool {
	return false
}

func (t *asserted) NoError() bool {
	return false
}

func (t *asserted) Equal() bool {
	return false
}

func (t *asserted) NotEqual() bool {
	return false
}
