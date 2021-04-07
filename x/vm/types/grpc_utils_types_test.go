package types

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

func TestVM_SdkIntToVmU128(t *testing.T) {
	// ok: 0x0
	{
		value := uint16(0x0)
		res, err := SdkIntToVmU128(sdk.NewInt(int64(value)))
		require.NoError(t, err)
		checkU128(t, value, res)
	}

	// ok: 0x1
	{
		value := uint16(0x1)
		res, err := SdkIntToVmU128(sdk.NewInt(int64(value)))
		require.NoError(t, err)
		checkU128(t, value, res)
	}

	// fail: < 0
	{
		_, err := SdkIntToVmU128(sdk.NewInt(-1))
		require.Error(t, err)
	}
}

func TestVM_SdkUintToVmU128(t *testing.T) {
	// ok: 0x0
	{
		value := uint16(0x0)
		res, err := SdkUintToVmU128(sdk.NewUint(uint64(value)))
		require.NoError(t, err)
		checkU128(t, value, res)
	}

	// ok: 0x1
	{
		value := uint16(0x1)
		res, err := SdkUintToVmU128(sdk.NewUint(uint64(value)))
		require.NoError(t, err)
		checkU128(t, value, res)
	}

	// ok: 0xFFF
	{
		value := uint16(0xFFF)
		res, err := SdkUintToVmU128(sdk.NewUint(uint64(value)))
		require.NoError(t, err)
		checkU128(t, value, res)
	}

	// ok: 0x1FFF
	{
		value := uint16(0x1FFF)
		res, err := SdkUintToVmU128(sdk.NewUint(uint64(value)))
		require.NoError(t, err)
		checkU128(t, value, res)
	}

	// fail: > 16 bytes
	{
		value := sdk.NewUintFromString("12345678901234567890123456789012345678901234567890")
		_, err := SdkUintToVmU128(value)
		require.Error(t, err)
	}
}

func TestVM_VmU128ToSdkInt(t *testing.T) {
	// nil
	{
		intValue := VmU128ToSdkInt(nil)

		require.EqualValues(t, 0, intValue.Int64())
		require.False(t, intValue.IsNegative())
	}

	// 1 byte
	{
		u128Value := &dvmTypes.U128{
			Buf: []byte{0xFF},
		}
		intValue := VmU128ToSdkInt(u128Value)

		require.EqualValues(t, 255, intValue.Int64())
		require.False(t, intValue.IsNegative())
	}

	// 4293844428 [0xFF, 0xEE, 0xDD, 0xCC]
	{
		u128Value := &dvmTypes.U128{
			Buf: []byte{0xCC, 0xDD, 0xEE, 0xFF},
		}
		intValue := VmU128ToSdkInt(u128Value)

		require.EqualValues(t, 4293844428, intValue.Int64())
		require.False(t, intValue.IsNegative())
	}
}

func checkU128(t *testing.T, inValue uint16, outValue *dvmTypes.U128) {
	require.NotNil(t, outValue)
	require.Len(t, outValue.Buf, 16)

	inValueBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(inValueBytes, inValue)
	for i := 0; i < 16; i++ {
		inValueByte := byte(0x0)
		if i < len(inValueBytes) {
			inValueByte = inValueBytes[i]
		}

		require.Equal(t, inValueByte, outValue.Buf[i], "in / out at index [%d]: %s / %s", i, hex.EncodeToString(inValueBytes), hex.EncodeToString(outValue.Buf))
	}
}
