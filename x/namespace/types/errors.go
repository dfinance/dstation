package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInternal       = sdkErrors.Register(ModuleName, 200, "internal")
	ErrNameDoesNotExist = sdkErrors.Register(ModuleName, 201, "name does not exist")
	ErrDomainNotHandlesByAddress = sdkErrors.Register(ModuleName, 202, "domain not handles by address")
)
