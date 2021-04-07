package types

import (
	"bytes"
	"fmt"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

const (
	ModuleName   = "vm"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	GovRouterKey = ModuleName
	//
	DelPoolName = "vm_delegation_pool"
)

var (
	// Storage keys
	KeyDelimiter = []byte(":")  // we should not rely on this delimiter for VMStorage (bytes.Split usage for instance) as VM accessPath.Path might include symbols like: [':', '@',..]
	VMKeyPrefix  = []byte("vm") // storage key prefix for VMStorage data
)

// GetVMStorageKey returns VMStorage key for dvmTypes.VMAccessPath.
func GetVMStorageKey(path *dvmTypes.VMAccessPath) []byte {
	if path == nil {
		return nil
	}

	return bytes.Join(
		[][]byte{
			path.Address,
			path.Path,
		},
		KeyDelimiter,
	)
}

// MustParseVMStorageKey parses VMStorage key and panics on failure.
func MustParseVMStorageKey(key []byte) *dvmTypes.VMAccessPath {
	// Key length is expected to be correct: {address_20bytes}:{path_at_least_1byte}
	expectedMinLen := VMAddressLength + len(KeyDelimiter) + 1
	if len(key) < expectedMinLen {
		panic(fmt.Errorf("VMKey (%s): invalid key length: expected / actual: %d / %d", string(key), expectedMinLen, len(key)))
	}

	// Calc indices
	addressStartIdx := 0
	addressEndIdx := addressStartIdx + VMAddressLength
	delimiterStartIdx := addressEndIdx
	delimiterEndIdx := delimiterStartIdx + len(KeyDelimiter)
	pathStartIdx := delimiterEndIdx

	// Split key
	addressValue := key[addressStartIdx:addressEndIdx]
	delimiterValue := key[delimiterStartIdx:delimiterEndIdx]
	pathValue := key[pathStartIdx:]

	// Validate
	if !bytes.Equal(delimiterValue, KeyDelimiter) {
		panic(fmt.Errorf("VMKey (%s): 1st delimiter value is invalid", string(key)))
	}
	if len(addressValue) < VMAddressLength {
		panic(fmt.Errorf("VMKey (%s): address length is invalid: expected / actual: %d / %d", string(key), VMAddressLength, len(addressValue)))
	}
	if len(pathValue) == 0 {
		panic(fmt.Errorf("VMKey (%s): path length is invalid: expected / actual: GT 1 / %d", string(key), len(pathValue)))
	}

	return &dvmTypes.VMAccessPath{
		Address: addressValue,
		Path:    pathValue,
	}
}
