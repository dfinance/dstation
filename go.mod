module github.com/dfinance/dstation

go 1.15

require (
	github.com/OneOfOne/xxhash v1.2.2
	github.com/atlassian/go-sentry-api v0.0.0-20210507210027-8175592dcf60
	github.com/cosmos/cosmos-sdk v0.42.6
	github.com/fsouza/go-dockerclient v1.7.2
	github.com/getsentry/sentry-go v0.11.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
