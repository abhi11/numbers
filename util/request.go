package util

import (
	"net/http"
	"net/url"
	"strings"
)

// Util method to fetch from an URL
func HttpGet(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Reads the response body and closes it. Then returns remaining Response object
// In case of errors, returns errors, and response(could be nil in err cases)
func HttpGetAndRead(url string, v interface{}) (*http.Response, error) {
	resp, err := HttpGet(url)
	if err != nil {
		return resp, err
	}
	err = Bind(resp.Body, v)
	return resp, err
}

// Get all the values for an url param; if no value is present
// return an empty list
func GetParamValues(r *http.Request, key string) []string {
	queryMap := r.URL.Query()
	vals, ok := queryMap[key]
	if !ok {
		// key doesn't exist
		vals = []string{}
	}
	return vals
}

// Checks if URL is valid
// Does the basic checkings, this makes the url selection faster
// As hitting the URLs are parallel virtually no time is wasted while
// calling an invalid url
func IsValidURL(str string) bool {
	length := len(str)
	if str == "" || length >= 2083 || strings.HasPrefix(str, ".") || !strings.HasPrefix(str, "http") {
		return false
	}

	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	if strings.HasPrefix(u.Host, ".") {
		return false
	}

	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return true
}

// Util to filter put all invalid urls from a
// list of strings
func GetValidURLs(strs []string) []string {
	validURLs := []string{}
	for _, v := range strs {
		if IsValidURL(v) {
			validURLs = append(validURLs, v)
		}
	}

	return validURLs
}
