package util

import (
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func ForceStringOf(t *testing.T, rc io.ReadCloser) string {
	s, err := StringOf(rc)
	require.NoError(t, err)
	return s
}
