package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInternal      = sdkErrors.Register(ModuleName, 100, "internal")
	ErrNotAuthorized = sdkErrors.Register(ModuleName, 101, "account is not a nominee")
)
