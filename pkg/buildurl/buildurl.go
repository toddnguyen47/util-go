package buildurl

import (
	url2 "net/url"
	"strconv"
	"strings"
)

// Build a URL based on a baseUrl, pathUrl, queryParams
func Build(baseUrl, pathUrl string, queryParams map[string]interface{}) string {
	url, err := url2.Parse(baseUrl)
	if err != nil {
		return ""
	}

	paths := strings.Split(url.Path, "/")
	var noEmptyStringPaths []string
	for _, path := range paths {
		trimmedPath := strings.TrimSpace(path)
		if len(trimmedPath) > 0 {
			noEmptyStringPaths = append(noEmptyStringPaths, trimmedPath)
		}
	}
	noEmptyStringPaths = append(noEmptyStringPaths, pathUrl)
	url.Path = strings.Join(noEmptyStringPaths, "/")

	query := url.Query()
	for key, val := range queryParams {
		query.Set(key, convertVal(val))
	}
	url.RawQuery = query.Encode()
	return url.String()
}

// convertVal - Convert interface to either string or int. Returns empty string by default
func convertVal(val interface{}) string {
	var s1 string
	if stringVal, ok := val.(string); ok {
		s1 = stringVal
	} else if intVal, ok := val.(int); ok {
		s1 = strconv.Itoa(intVal)
	}

	return s1
}
