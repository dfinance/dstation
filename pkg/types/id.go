package types

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.CustomProtobufType = ID{}

type ID sdk.Uint

func (id ID) uint() sdk.Uint { return sdk.Uint(id) }

func (id ID) UInt64() uint64 { return id.uint().Uint64() }

func (id ID) Valid() error {
	nilID := ID(sdk.Uint{})
	if reflect.DeepEqual(id, nilID) {
		return fmt.Errorf("nil")
	}

	return nil
}

func (id ID) Equal(id2 ID) bool { return id.uint().Equal(id2.uint()) }

func (id ID) LT(id2 ID) bool { return id.uint().LT(id2.uint()) }

func (id ID) LTE(id2 ID) bool { return id.uint().LTE(id2.uint()) }

func (id ID) GT(id2 ID) bool { return id.uint().GT(id2.uint()) }

func (id ID) GTE(id2 ID) bool { return id.uint().GTE(id2.uint()) }

func (id ID) Incr() ID { return ID(id.uint().Incr()) }

func (id ID) Decr() ID { return ID(id.uint().Decr()) }

func (id ID) String() string { return id.uint().String() }

func (id ID) Marshal() ([]byte, error) {
	return id.uint().Marshal()
}

func (id ID) MarshalTo(data []byte) (n int, err error) {
	u := id.uint()
	return u.MarshalTo(data)
}

func (id ID) Unmarshal(data []byte) error {
	u := id.uint()
	return u.Unmarshal(data)
}

func (id ID) Size() int {
	u := id.uint()
	return u.Size()
}

func (id ID) MarshalJSON() ([]byte, error) {
	u := id.uint()
	return u.MarshalJSON()
}

func (id ID) UnmarshalJSON(data []byte) error {
	u := id.uint()
	return u.UnmarshalJSON(data)
}

func NewZeroID() ID { return ID(sdk.ZeroUint()) }

func NewIDFromUint64(id uint64) ID { return ID(sdk.NewUint(id)) }

func NewIDFromString(str string) (retID ID, retErr error) {
	// sdk.NewUintFromString might panic
	defer func() {
		if r := recover(); r != nil {
			retErr = fmt.Errorf("%q cannot be converted to big.Int", str)
		}
	}()

	if str == "" {
		return ID{}, fmt.Errorf("empty")
	}

	return ID(sdk.NewUintFromString(str)), nil
}
