package types

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVM_MustParsePathKey(t *testing.T) {
	address := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	path := []byte{1, 2, 4, 8, 16, 32, 64, 128, 255}

	// ok
	{
		key := bytes.Join(
			[][]byte{
				address,
				path,
			}, KeyDelimiter)

		accessPath := MustParseVMStorageKey(key)
		require.EqualValues(t, address, accessPath.Address)
		require.EqualValues(t, path, accessPath.Path)
	}

	// fail: wrong address length
	{
		key := bytes.Join(
			[][]byte{
				address[:len(address)-2],
				path,
			}, KeyDelimiter)

		require.Panics(t, func() {
			MustParseVMStorageKey(key)
		})
	}

	// fail: empty path
	{
		key := bytes.Join(
			[][]byte{
				address,
				{},
			}, KeyDelimiter)

		require.Panics(t, func() {
			MustParseVMStorageKey(key)
		})
	}

	// fail: wrong delimiter
	{
		key := bytes.Join(
			[][]byte{
				address,
				{},
			}, []byte("@"))

		require.Panics(t, func() {
			MustParseVMStorageKey(key)
		})
	}
}
