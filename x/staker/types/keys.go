package types

const (
	ModuleName   = "staker"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	// Storage keys
	CallsPrefix = []byte("calls")      // storage key prefix for storing Call data
	LastCallId  = []byte("lastCallId") // storage key for storing the latest Call ID
)
