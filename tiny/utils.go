package tiny

import "strings"

type pathPattern string
type HandlerFunc func(*Context)
type MiddlewareFunc func(HandlerFunc) HandlerFunc

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

func getPathParamKeys(path string) ([]string, error) {
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
