package util

import (
	"io"
	"strings"
)

func StringOf(body io.ReadCloser) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, body)

	if err != nil {
		return "", err
	}

	s := buf.String()
	return s, err
}
