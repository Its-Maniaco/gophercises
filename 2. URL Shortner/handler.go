package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if pathsToUrls[r.URL.Path] == "" {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusPermanentRedirect)
		}
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// This pathsToUrl is same as what we did in MapHandler after Unmarshaling
	pathsToUrl, err := yamlParse(yml)
	handlerMap := mapForYAML(pathsToUrl)

	return MapHandler(handlerMap, fallback), err
}

func yamlParse(yml []byte) ([]yamlStruct, error) {
	var pathsToUrls []yamlStruct
	err := yaml.Unmarshal(yml, &pathsToUrls)
	return pathsToUrls, err
}

type yamlStruct struct {
	path string //if this path
	url  string // then this url
}

// creating map from the Yaml to directly use in MapHandler
func mapForYAML(pathsToURLs []yamlStruct) map[string]string {
	yamlMap := make(map[string]string, len(pathsToURLs))
	for _, path := range pathsToURLs {
		yamlMap[path.path] = path.url
	}
	return yamlMap
}
