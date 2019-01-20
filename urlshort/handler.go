package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	// How do I get the current path?

	// In the examples the object passed to ListenAndServe is
	// a Handler. However, this functions demands a HandlerFunc
	// as a return value.
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Interesting, as soon as you write here to the response writer
		// the redirect will not work and instead print an anchor
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type Pair struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func ParseYML(yml []byte) ([]Pair, error) {

	var result []Pair
	err := yaml.Unmarshal(yml, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func BuildMap(data []Pair) (map[string]string, error) {
	result := make(map[string]string)
	for _, pair := range data {
		result[pair.Path] = pair.Url
	}
	return result, nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYML, err := ParseYML(yml)

	if err != nil {
		return nil, err
	}

	pathMap, err := BuildMap(parsedYML)

	if err != nil {
		return nil, err
	}

	return MapHandler(pathMap, fallback), nil
}
