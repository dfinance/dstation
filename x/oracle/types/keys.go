package types

import (
	"bytes"
	"encoding/binary"

	"github.com/dfinance/dstation/pkg/types"
)

const (
	ModuleName   = "oracle"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	// Storage keys
	KeyDelimiter    = []byte(":")
	AssetsPrefix    = []byte("assets")    // storage key prefix for storing Asset data
	OraclesPrefix   = []byte("oracles")   // storage key prefix for storing Oracle data
	CurPricesPrefix = []byte("curPrices") // storage key prefix for storing CurrentPrice data
	RawPricesPrefix = []byte("rawPrices") // storage key prefix for storing RawPrice data
)

// GetOraclesKey returns OraclesPrefix storage key.
func GetOraclesKey(address string) []byte {
	return []byte(address)
}

// GetRawPricesKey returns RawPricesPrefix storage key.
func GetRawPricesKey(assetCode types.AssetCode, blockHeight int64, oracleAddress string) []byte {
	assetCodeBz, _ := assetCode.Marshal()

	blockHeightBz := make([]byte, 8)
	binary.LittleEndian.PutUint64(blockHeightBz, uint64(blockHeight))

	return bytes.Join(
		[][]byte{
			assetCodeBz,
			blockHeightBz,
			[]byte(oracleAddress),
		},
		KeyDelimiter,
	)
}

// GetRawPricesKeyPrefix returns RawPricesPrefix storage key prefix without oracleAddress.
func GetRawPricesKeyPrefix(assetCode types.AssetCode, blockHeight int64) []byte {
	assetCodeBz, _ := assetCode.Marshal()

	blockHeightBz := make([]byte, 8)
	binary.LittleEndian.PutUint64(blockHeightBz, uint64(blockHeight))

	return bytes.Join(
		[][]byte{
			assetCodeBz,
			blockHeightBz,
		},
		KeyDelimiter,
	)
}
