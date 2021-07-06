package types

const (
	ModuleName   = "namespace"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	WhoisPrefix     = []byte("whois")      // storage key prefix for storing Call data by ID
	LastWhoisId      = []byte("lastWhoisId") // storage key for storing the latest Call ID
)
