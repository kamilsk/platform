package testing

import "testing"

// Assertable makes testing.T assertable.
func Assertable(t *testing.T) *assertable {
	return (*assertable)(t)
}

type assertable testing.T

func (t *assertable) True() bool {
	return false
}

func (t *assertable) False() bool {
	return false
}

func (t *assertable) Error() bool {
	return false
}

func (t *assertable) EqualError() bool {
	return false
}

func (t *assertable) MatchError() bool {
	return false
}

func (t *assertable) NoError() bool {
	return false
}

func (t *assertable) Equal() bool {
	return false
}

func (t *assertable) NotEqual() bool {
	return false
}

// Strict makes testing.T strictly assertable.
func Strict(t *testing.T) *strictable {
	return (*strictable)(t)
}

type strictable testing.T

func (t *strictable) True() {
}

func (t *strictable) False() {
}

func (t *strictable) Error() {
}

func (t *strictable) EqualError() {
}

func (t *strictable) MatchError() {
}

func (t *strictable) NoError() {
}

func (t *strictable) Equal() {
}

func (t *strictable) NotEqual() {
}
