package util

import (
	"encoding/json"
	"io"
)

// Bind reads the body into an interface and closes the body
func Bind(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	err := json.NewDecoder(body).Decode(v)
	return err
}

// Takes an interface and converts it into array of bytes
func ToJSONBytes(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
