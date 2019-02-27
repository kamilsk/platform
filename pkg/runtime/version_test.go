package runtime_test

import (
	"testing"

	. "github.com/kamilsk/platform/pkg/runtime"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	version := Version()
	assert.NotEmpty(t, version.Major)
	assert.NotEmpty(t, version.Minor)
}

func TestVersion_Compare(t *testing.T) {
	tests := []struct {
		name    string
		target  GoVersion
		compare func(GoVersion, GoVersion) bool
	}{
		{"before", GoVersion{Major: 2}, GoVersion.Before},
		{"later", GoVersion{Major: 1}, GoVersion.Later},
		{"much later", GoVersion{}, GoVersion.Later},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.True(t, tc.compare(Version(), tc.target))
		})
	}
	t.Run("full comparison", func(t *testing.T) {
		base := GoVersion{Major: 1, Minor: 9}
		v1, v2, v3 := base, base, base
		v1.Patch, v2.Patch, v3.Patch = 1, 2, 3
		assert.True(t, v1.Before(v2) && v2.Before(v3))
		assert.True(t, v3.Later(v2) && v2.Later(v1))
		assert.True(t, v2.Later(v1) && v2.Before(v3))
		assert.True(t, !base.Later(base) && !base.Before(base))
	})
}
