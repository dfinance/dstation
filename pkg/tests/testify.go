package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func CheckPanicErrorContains(t *testing.T, handler func(), errSubStr string) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "handler did not panic")

		err, ok := r.(error)
		require.True(t, ok, "panic obj is not an error")

		require.Contains(t, err.Error(), errSubStr)
	}()

	handler()
}
