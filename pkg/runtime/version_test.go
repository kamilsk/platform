package runtime_test

import (
	"strings"
	"testing"
	"time"

	. "github.com/kamilsk/platform/pkg/runtime"
	"github.com/stretchr/testify/assert"
)

func TestVersion_Compare(t *testing.T) {
	version := Version()
	if unstable(version.Raw) {
		version.Major, version.Minor, version.Patch, version.Raw = 1, 12, 0, "go1.12"
	}

	tests := []struct {
		name    string
		target  GoVersion
		compare func(GoVersion, GoVersion) bool
	}{
		{"before", GoVersion{Major: 2}, GoVersion.Before},
		{"later", GoVersion{Major: 1}, GoVersion.Later},
		{"much later", GoVersion{}, GoVersion.Later},
		{"equal", version, GoVersion.Equal},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.True(t, tc.compare(version, tc.target))
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

func ahead(current GoVersion, target struct {
	version GoVersion
	release string
}) bool {
	if current.Equal(target.version) {
		return true
	}
	// devel +61170f85e6 Thu Feb 28 00:24:56 2019 +0000
	if !unstable(current.Raw) {
		return current.Later(target.version)
	}
	prefix := "devel +61170f85e6 "
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	release, _ := time.Parse(layout, target.release)
	control, _ := time.Parse(layout, current.Raw[len(prefix):])
	return control.After(release)
}

func unstable(version string) bool {
	return strings.HasPrefix(version, "devel")
}

var go112 = struct {
	version GoVersion
	release string
}{
	version: GoVersion{Major: 1, Minor: 12, Raw: "go1.12"},
	release: "Mon Feb 25 16:47:57 2019 -0500",
}
