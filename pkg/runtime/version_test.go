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

func ahead(version GoVersion, timestamp string) bool {
	// devel +61170f85e6 Thu Feb 28 00:24:56 2019 +0000
	if !unstable(version.Raw) {
		return false
	}
	prefix := "devel +61170f85e6 "
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	target, _ := time.Parse(layout, timestamp)
	current, _ := time.Parse(layout, version.Raw[len(prefix):])
	return current.After(target)
}

func unstable(version string) bool {
	return strings.HasPrefix("devel", version)
}

var go112 = struct {
	version GoVersion
	release string
}{
	version: GoVersion{Major: 1, Minor: 12},
	release: "Mon Feb 25 16:47:57 2019 -0500",
}
