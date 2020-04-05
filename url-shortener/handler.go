package urlshortener

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToURLs[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type pathURLMapping struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathURLs []pathURLMapping
	if err := yaml.Unmarshal(yamlBytes, &pathURLs); err != nil {
		return nil, err
	}

	var pathToURLS map[string]string
	for _, mapping := range pathURLs {
		pathToURLS[mapping.Path] = mapping.URL
	}

	return MapHandler(pathToURLS, fallback), nil
}
