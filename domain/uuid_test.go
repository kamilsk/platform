package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/domain"
)

func TestUUID(t *testing.T) {
	tests := []struct {
		name   string
		uuid   UUID
		assert func(assert.TestingT, bool, ...interface{}) bool
	}{
		{"empty", "", assert.False},
		{"invalid", "abc-def-ghi", assert.False},
		{"not v4 [0]", "41ca5e09-3ce2-0094-b108-3ecc257c6fa4", assert.False},
		{"not v4 [1]", "41ca5e09-3ce2-1094-b108-3ecc257c6fa4", assert.False},
		{"not v4 [2]", "41ca5e09-3ce2-2094-b108-3ecc257c6fa4", assert.False},
		{"not v4 [3]", "41ca5e09-3ce2-3094-b108-3ecc257c6fa4", assert.False},
		{"v4 [lower]", "41ca5e09-3ce2-4094-b108-3ecc257c6fa4", assert.True},
		{"v4 [upper]", "41CA5E09-3CE2-4094-B108-3ECC257C6FA4", assert.True},
		{"not v4 [5]", "41ca5e09-3ce2-5094-b108-3ecc257c6fa4", assert.False},
	}
	for _, test := range tests {
		assert.Equal(t, test.uuid == "", test.uuid.IsEmpty(), test.name)
		assert.Equal(t, test.uuid, UUID(test.uuid.String()), test.name)
		test.assert(t, test.uuid.IsValid())
	}
}
