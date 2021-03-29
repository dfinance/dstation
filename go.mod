module github.com/dfinance/dstation

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.2
	github.com/dfinance/dvm-proto/go v0.0.0-20201007122036-27be7297df4e
	github.com/dfinance/glav v0.0.0-20200814081332-c4701f6c12a6
	github.com/dfinance/lcs v0.1.7-big
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
