package apitools

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/blang/semver/v4"
)

var errAPINotFound = errors.New("api not found")

var (
	typeMap   = map[string]map[string]typeRef{}
	typeMapMu = &sync.RWMutex{}
)

type typeRef struct {
	Type  reflect.Type
	Hooks []func(interface{})
}

func (ref typeRef) new() interface{} {
	typ := reflect.New(ref.Type).Interface()
	for _, hook := range ref.Hooks {
		hook(typ)
	}
	return typ
}

// ResolveOption a customization on how types are resolved
type ResolveOption interface {
	apply(ref *typeRef)
}

// WithAlias Option allows a type to be resolved by names other than its type
// name
func WithAlias(alias ...string) ResolveOption {
	return aliasOpt{
		Aliases: alias,
	}
}

type aliasOpt struct {
	Aliases []string
}

// Noop to fit Option interface
func (aliasOpt) apply(*typeRef) {}

// WithResolveHook allows modules to preform initalization on resolved types
func WithResolveHook(fn func(interface{})) ResolveOption {
	return hookOpt(fn)
}

type hookOpt func(interface{})

func (fn hookOpt) apply(ref *typeRef) {
	ref.Hooks = append(ref.Hooks, fn)
}

// RegisterType allows modules to register API types to be resolved.
func RegisterType(apiGroup string, t interface{}, opts ...ResolveOption) {
	var typeAliases []string

	rrt := reflect.ValueOf(t)
	rType := reflect.Indirect(rrt).Type()
	ref := typeRef{Type: rType}
	for _, opt := range opts {
		if alias, ok := opt.(aliasOpt); ok {
			typeAliases = append(typeAliases, alias.Aliases...)
		}
		opt.apply(&ref)
	}
	typeMapMu.Lock()
	defer typeMapMu.Unlock()
	if _, ok := typeMap[apiGroup]; !ok {
		typeMap[apiGroup] = map[string]typeRef{}
	}
	typeMap[apiGroup][rType.Name()] = ref
	for _, alias := range typeAliases {
		typeMap[apiGroup][alias] = ref
	}
}

// Resolve resolves the raw type for the requested api version and type.
func Resolve(apiVersion string, typename string) (interface{}, error) {

	// Guard read access to packageMap
	typeMapMu.RLock()
	defer typeMapMu.RUnlock()
	apiGroup, reqVer := parseAPIVersion(apiVersion)

	group, ok := typeMap[apiGroup]
	if !ok {
		return nil, fmt.Errorf("api group %s has not been registered", apiGroup)
	}
	ref, ok := group[typename]
	if !ok {
		return nil, errAPINotFound
	}

	if foundVer, err := versionOf(ref.Type); err == nil {
		if semverGreater(reqVer, foundVer) {
			return nil, fmt.Errorf("requested version was %s, but only %s is available", reqVer, foundVer)
		}
	}

	return ref.new(), nil
}

func semverGreater(s1, s2 string) bool {
	s1Ver, err := semver.ParseTolerant(s1)
	if err != nil {
		// semver should be validated before being passed here
		return false
	}
	s2Ver, err := semver.ParseTolerant(s2)
	if err != nil {
		// semver should be validated before being passed here
		return false
	}
	return s1Ver.GT(s2Ver)
}
