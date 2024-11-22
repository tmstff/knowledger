package util

import (
	"io"
	"strings"
)

func StringOf(rc io.ReadCloser) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, rc)

	if err != nil {
		return "", err
	}

	s := buf.String()
	return s, err
}
