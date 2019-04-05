package testing

import "testing"

// Required makes testing.T required.
func Required(t *testing.T) *required {
	return (*required)(t)
}

type required testing.T

func (t *required) True() {
}

func (t *required) False() {
}

func (t *required) Error() {
}

func (t *required) EqualError() {
}

func (t *required) MatchError() {
}

func (t *required) NoError() {
}

func (t *required) Equal() {
}

func (t *required) NotEqual() {
}
