package tiny

import "strings"

type pathPattern string
type pathParamKeys []string

func getPathPattern(path string) (pathPattern, error) {
	routeSlices := strings.Split(path, "/")
	for i := range routeSlices {
		slice := routeSlices[i]
		if len(slice) == 0 {
			continue
		}
		if slice[0] == ':' {
			routeSlices[i] = "(.+)"
		}
	}

	return pathPattern(strings.Join(routeSlices, "/")), nil
}

func getPathParamKeys(path string) (pathParamKeys, error) {
	var urlParamKeys []string = make([]string, 0)
	routeSlices := strings.Split(path, "/")
	for i := range routeSlices {
		slice := routeSlices[i]
		if len(slice) == 0 {
			continue
		}
		if slice[0] == ':' {
			routeSlices[i] = "(.+)"
			urlParamKeys = append(urlParamKeys, slice[1:])
		}
	}

	return urlParamKeys, nil
}
