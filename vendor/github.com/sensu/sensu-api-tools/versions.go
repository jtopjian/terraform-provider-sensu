package apitools

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"runtime/debug"
	"sort"
	"strings"
	"sync"

	"github.com/blang/semver/v4"
)

var (
	errBuildInfoUnavailable = errors.New("build info unavailable")
	deps                    []*debug.Module
	depsSetup               sync.Once
)

// APIModuleVersions returns a map of Sensu API modules that are compiled into
// the product.
func APIModuleVersions() map[string]string {
	apiModuleVersions := make(map[string]string)
	typeMapMu.Lock()
	defer typeMapMu.Unlock()
	for groupName, group := range typeMap {
		var first reflect.Type
		for _, typ := range group {
			first = typ.Type
			break
		}
		groupVersion, err := versionOf(first)
		if err != nil {
			_, groupVersion = parseAPIVersion(groupName)
		}
		apiModuleVersions[groupName] = groupVersion
	}
	return apiModuleVersions
}

// versionOf interrogates the build environment for the version of the module
// defining a type.
func versionOf(typ reflect.Type) (string, error) {
	packagePath := typ.PkgPath()
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok || buildInfo.Deps == nil {
		return "", errBuildInfoUnavailable
	}
	for _, mod := range buildInfo.Deps {
		if strings.HasPrefix(packagePath, mod.Path) {
			return mod.Version, nil
		}
	}
	return "", fmt.Errorf("error finding build dependency for type %s", typ)
}

func buildDependencies() []*debug.Module {
	depsSetup.Do(func() {
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok || buildInfo.Deps == nil {
			return
		}
		for _, dep := range buildInfo.Deps {
			deps = append(deps, dep)
		}
		// sort by module path length descending so that when matching package
		// names to a module we find the more specific modules first.
		// e.g. github.com/sensu/sensu-go/api/core/v2 before github.com/sensu/sensu-go
		sort.Slice(deps, func(i, j int) bool {
			return len(deps[i].Path) > len(deps[j].Path)
		})
	})
	return deps
}

// parseAPIVersion parses an api_version that looks like the following:
//
// core/v2
// core/v2.2
// core/v2.2.1
//
// It returns the name of the apiGroup (core/v2), and the semantic version
// (v2.0.0, v2.2.0, v2.2.1). A leading 'v' is included, keeping with how Go
// modules express their versions.
//
// If ParseAPIVersion can't determine the version, for instance if it's passed
// a string that does not seem to be a versioned API group, it will return its
// input as the apiGroup, and v0.0.0 as the version.
func parseAPIVersion(apiVersion string) (apiGroup, semVer string) {
	group, version := path.Split(apiVersion)
	if version == "" {
		// There is no version for the API group, which is fine.
		return group, "v0.0.0"
	}
	semver, err := semver.ParseTolerant(version)
	if err != nil {
		// It's not the expected format
		return apiVersion, "v0.0.0"
	}
	apiGroup = path.Join(group, fmt.Sprintf("v%d", semver.Major))
	semVer = fmt.Sprintf("v%d.%d.%d", semver.Major, semver.Minor, semver.Patch)
	return apiGroup, semVer
}
