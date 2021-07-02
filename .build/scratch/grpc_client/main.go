package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	gRpcUrl = "grpc.demo1.dfinance.co:443"
	//gRpcUrl = "127.0.0.1:9091"
)

func main() {
	sdkConfig := sdk.GetConfig()
	dnConfig.SetConfigBech32Prefixes(sdkConfig)
	sdkConfig.Seal()

	accAddr, err := sdk.AccAddressFromBech32("wallet1zlvupgunlyad0kswrg9y06y75qk5cgkjm6r8ls")
	if err != nil {
		panic(fmt.Errorf("sdk.AccAddressFromBech32: %w", err))
	}

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		panic(fmt.Errorf("x509.SystemCertPool: %w", err))
	}
	tlsConfig := &tls.Config{
		RootCAs: rootCAs,
	}

	dialOpts := []grpc.DialOption{
		//grpc.WithInsecure(),
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}
	clientCon, err := grpc.Dial(gRpcUrl, dialOpts...)
	if err != nil {
		panic(fmt.Errorf("pkg.GetGRpcClientConnection: %w", err))
	}
	defer clientCon.Close()

	client := bankTypes.NewQueryClient(clientCon)

	resp, err := client.AllBalances(context.Background(), &bankTypes.QueryAllBalancesRequest{
		Address: accAddr.String(),
	})
	if err != nil {
		panic(fmt.Errorf("client.AllBalances: %w", err))
	}

	fmt.Printf("Addr (%s): Coins: %s\n", accAddr, resp.Balances)
}
