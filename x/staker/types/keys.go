package types

const (
	ModuleName   = "staker"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	// Storage keys
	CallsPrefix     = []byte("calls")      // storage key prefix for storing Call data by ID
	UniqueIdsPrefix = []byte("uniqueIds")  // storage key prefix to store Call UniqueId <-> Id match
	LastCallId      = []byte("lastCallId") // storage key for storing the latest Call ID
)
