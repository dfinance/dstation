package types

import (
	"bytes"
	"fmt"

	"github.com/dfinance/dvm-proto/go/vm_grpc"
)

const (
	ModuleName   = "vm"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	GovRouterKey = ModuleName
)

var (
	// Storage keys
	KeyDelimiter = []byte(":")  // we should rely on this delimiter (for bytes.Split for example) as VM accessPath.Path might include symbols like: [':', '@',..]
	VMKey        = []byte("vm") // storage key prefix for VMStorage data
)

// GetVMStorageKey returns VMStorage key for vm_grpc.VMAccessPath.
func GetVMStorageKey(path *vm_grpc.VMAccessPath) []byte {
	if path == nil {
		return nil
	}

	return bytes.Join(
		[][]byte{
			VMKey,
			path.Address,
			path.Path,
		},
		KeyDelimiter,
	)
}

// GetVMStorageKeyPrefix returns VMStorage keys prefix (used for iteration).
func GetVMStorageKeyPrefix() []byte {
	return append(VMKey, KeyDelimiter...)
}

// MustParseVMStorageKey parses VMStorage key and panics on failure.
func MustParseVMStorageKey(key []byte) *vm_grpc.VMAccessPath {
	accessPath := vm_grpc.VMAccessPath{}

	// we expect key to be correct: vm:{address_20bytes}:{path_at_least_1byte}
	expectedMinLen := len(VMKey) + len(KeyDelimiter) + VMAddressLength + len(KeyDelimiter) + 1
	if len(key) < expectedMinLen {
		panic(fmt.Errorf("key %q: invalid length: min expected: %d", string(key), expectedMinLen))
	}

	// calc indices (end index is the next one of the real end idx)
	prefixStartIdx := 0
	prefixEndIdx := prefixStartIdx + len(VMKey)
	delimiterFirstStartIdx := prefixEndIdx
	delimiterFirstEndIdx := delimiterFirstStartIdx + len(KeyDelimiter)
	addressStartIdx := delimiterFirstEndIdx
	addressEndIdx := addressStartIdx + VMAddressLength
	delimiterSecondStartIdx := addressEndIdx
	delimiterSecondEndIdx := delimiterSecondStartIdx + len(KeyDelimiter)
	pathStartIdx := delimiterSecondEndIdx

	// split key
	prefixValue := key[prefixStartIdx:prefixEndIdx]
	delimiterFirstValue := key[delimiterFirstStartIdx:delimiterFirstEndIdx]
	addressValue := key[addressStartIdx:addressEndIdx]
	delimiterSecondValue := key[delimiterSecondStartIdx:delimiterSecondEndIdx]
	pathValue := key[pathStartIdx:]

	// validate
	if !bytes.Equal(prefixValue, VMKey) {
		panic(fmt.Errorf("key %q: prefix: invalid", string(key)))
	}
	if !bytes.Equal(delimiterFirstValue, KeyDelimiter) {
		panic(fmt.Errorf("key %q: 1st delimiter: invalid", string(key)))
	}
	if !bytes.Equal(delimiterSecondValue, KeyDelimiter) {
		panic(fmt.Errorf("key %q: 2nd delimiter: invalid", string(key)))
	}

	accessPath.Address = addressValue
	accessPath.Path = pathValue

	return &accessPath
}
