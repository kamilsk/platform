package runtime

import (
	"runtime"
	"strconv"
	"strings"
)

// Version returns information about the current Go version.
func Version() GoVersion {
	var version = GoVersion{Raw: runtime.Version()}
	if strings.HasPrefix(version.Raw, "go") {
		divided := strings.Split(strings.TrimPrefix(version.Raw, "go"), ".")
		if len(divided) > 2 {
			divided = divided[:3]
		}
		converted := make([]int, 3)
		for i := range divided {
			converted[i], _ = strconv.Atoi(divided[i])
		}
		version.Major, version.Minor, version.Patch = converted[0], converted[1], converted[2]
	}
	return version
}

// GoVersion holds information about a Go version.
type GoVersion struct {
	Major int
	Minor int
	Patch int
	Raw   string
}

// Before returns true if the current Go version is before
// then the target Go version.
func (current GoVersion) Before(target GoVersion) bool {
	return current.compare(target) == -1
}

// Later returns true if the current Go version is later
// then the target Go version.
func (current GoVersion) Later(target GoVersion) bool {
	return current.compare(target) == 1
}

func (current GoVersion) compare(with GoVersion) int {
	if current.Major != with.Major {
		if current.Major > with.Major {
			return 1
		}
		return -1
	}
	if current.Minor != with.Minor {
		if current.Minor > with.Minor {
			return 1
		}
		return -1
	}
	if current.Patch != with.Patch {
		if current.Patch > with.Patch {
			return 1
		}
		return -1
	}
	return 0
}
