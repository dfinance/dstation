package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInternal         = sdkErrors.Register(ModuleName, 100, "internal")
	ErrNotAuthorized    = sdkErrors.Register(ModuleName, 101, "account is not a nominee")
	ErrOracleNotFound   = sdkErrors.Register(ModuleName, 102, "Oracle not registered")
	ErrInvalidPostPrice = sdkErrors.Register(ModuleName, 103, "invalid postPrice values")
)
