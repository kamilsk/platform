package testing

import (
	"fmt"
	"testing"
)

// Asserted makes testing.T asserted.
func Asserted(t *testing.T) *asserted {
	return (*asserted)(t)
}

type asserted testing.T

func (t *asserted) Cause(err error) error {
	for err != nil {
		cause, ok := err.(interface {
			Cause() error
		})
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

func (t *asserted) Equal() bool {
	return false
}

func (t *asserted) NotEqual() bool {
	return false
}

func (t *asserted) True(value bool, message ...interface{}) bool {
	if !value {
		t.Error(append([]interface{}{"should be true"}, message...)...)
	}

	return value
}

func (t *asserted) False(value bool, message ...interface{}) bool {
	if value {
		t.Error(append([]interface{}{"should be false"}, message...)...)
	}

	return !value
}

func (t *asserted) EqualError(err error, expected string, message ...interface{}) bool {
	if !t.HasError(err, message...) {
		return false
	}

	//

	return false
}

func (t *asserted) HasError(err error, message ...interface{}) bool {
	if err == nil {
		t.Error(append([]interface{}{"an error is expected but got nil"}, message...)...)
	}

	return false
}

func (t *asserted) MatchError() bool {
	return false
}

func (t *asserted) NoError(err error, message ...interface{}) bool {
	if err != nil {
		t.Error(append([]interface{}{fmt.Sprintf("an unexpected error: %+v", err)}, message...)...)
	}

	return false
}
