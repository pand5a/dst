package guess

import (
	"regexp"
	"strings"
)

func New() RestorerResolver {
	return RestorerResolver{}
}

func WithMap(m map[string]string) RestorerResolver {
	return RestorerResolver(m)
}

// RestorerResolver is a map of package path -> package name. Names are resolved from this map, and
// if a name doesn't exist in the map, the package name is guessed from the last part of the path
// (after the last slash).
type RestorerResolver map[string]string

func (r RestorerResolver) ResolvePackage(importPath string) (string, error) {
	if n, ok := r[importPath]; ok {
		return n, nil
	}
	if !strings.Contains(importPath, "/") {
		return importPath, nil
	}

	newPath := importPath

	ridx := strings.LastIndex(importPath, "/")
	suffix := importPath[ridx:]
	ok, _ := regexp.MatchString("^/v[0-9]*$", suffix)
	if ok {
		newPath = importPath[:ridx]
	}

	return newPath[strings.LastIndex(newPath, "/")+1:], nil
}
