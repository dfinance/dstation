package types

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

// SdkIntToVmU128 converts sdk.Int to dvmTypes.U128.
func SdkIntToVmU128(value sdk.Int) (*dvmTypes.U128, error) {
	if value.IsNegative() {
		return nil, fmt.Errorf("sdk.Int is negative (not Uint): %s", value.String())
	}

	return SdkUintToVmU128(sdk.NewUintFromBigInt(value.BigInt()))
}

// SdkUintToVmU128 converts sdk.Uint to dvmTypes.U128.
func SdkUintToVmU128(value sdk.Uint) (*dvmTypes.U128, error) {
	if value.BigInt().BitLen() > 128 {
		return nil, fmt.Errorf("invalid bitLen %d", value.BigInt().BitLen())
	}

	// BigInt().Bytes() returns BigEndian format, reverse it
	valueBytes := value.BigInt().Bytes()
	for left, right := 0, len(valueBytes)-1; left < right; left, right = left+1, right-1 {
		valueBytes[left], valueBytes[right] = valueBytes[right], valueBytes[left]
	}

	// Extend to 16 bytes
	if len(valueBytes) < 16 {
		zeros := make([]byte, 16-len(valueBytes))
		valueBytes = append(valueBytes, zeros...)
	}

	return &dvmTypes.U128{Buf: valueBytes}, nil
}

// VmU128ToSdkInt converts dvmTypes.U128 to sdk.Int.
func VmU128ToSdkInt(value *dvmTypes.U128) sdk.Int {
	if value == nil || len(value.Buf) == 0 {
		return sdk.ZeroInt()
	}

	// BigInt is BigEndian, convert U128 Little to Big
	for left, right := 0, len(value.Buf)-1; left < right; left, right = left+1, right-1 {
		value.Buf[left], value.Buf[right] = value.Buf[right], value.Buf[left]
	}

	// New big.Int
	bigValue := big.NewInt(0)
	bigValue.SetBytes(value.Buf)

	return sdk.NewIntFromBigInt(bigValue)
}
