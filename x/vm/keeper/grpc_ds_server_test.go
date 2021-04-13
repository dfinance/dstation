package keeper_test

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/tests"
	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *KeeperMockVmTestSuite) TestDSServer_GetRaw() {
	ctx, keeper := s.ctx, s.keeper

	// ok: non-existing
	{
		vmPath := mock.GetRandomVMAccessPath()

		s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &dvmTypes.DSAccessPath{Address: vmPath.Address, Path: vmPath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvmTypes.ErrorCode_NO_DATA, resp.ErrorCode)
			s.Require().NotEmpty(resp.ErrorMessage)
			s.Require().Nil(resp.Blob)
		})
	}

	// ok
	{
		vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(8)
		keeper.SetValue(ctx, vmPath, writeSetData)

		s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &dvmTypes.DSAccessPath{Address: vmPath.Address, Path: vmPath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvmTypes.ErrorCode_NONE, resp.ErrorCode)
			s.Require().Empty(resp.ErrorMessage)
			s.Require().Equal(writeSetData, resp.Blob)
		})
	}
}

func (s *KeeperMockVmTestSuite) TestDSServer_GetCurrencyInfo() {
	ctx, bankKeeper := s.ctx, s.app.DnApp.BankKeeper

	// fail: input
	{
		// denom
		{
			s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
				resp, err := client.GetCurrencyInfo(ctx, &dvmTypes.CurrencyInfoRequest{Ticker: ""})
				s.Require().NoError(err)
				s.Require().NotNil(resp)

				s.Require().Equal(dvmTypes.ErrorCode_BAD_REQUEST, resp.ErrorCode)
				s.Require().NotEmpty(resp.ErrorMessage)
			})
		}
	}

	// ok: non-existing
	{
		s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
			resp, err := client.GetCurrencyInfo(ctx, &dvmTypes.CurrencyInfoRequest{Ticker: "abc"})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvmTypes.ErrorCode_NO_DATA, resp.ErrorCode)
			s.Require().NotEmpty(resp.ErrorMessage)
			s.Require().Nil(resp.Info)
		})
	}

	// ok: xfi
	{
		s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))
		curTotalSupply := bankKeeper.GetSupply(ctx).GetTotal().AmountOf(dnConfig.MainDenom)

		curTotalSupplyU128, err := types.SdkIntToVmU128(curTotalSupply)
		s.Require().NoError(err)

		s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
			resp, err := client.GetCurrencyInfo(ctx, &dvmTypes.CurrencyInfoRequest{Ticker: "xfi"})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvmTypes.ErrorCode_NONE, resp.ErrorCode)
			s.Require().Empty(resp.ErrorMessage)

			s.Require().Equal([]byte("xfi"), resp.Info.Denom)
			s.Require().EqualValues(18, resp.Info.Decimals)
			s.Require().EqualValues(curTotalSupplyU128.Buf, resp.Info.TotalSupply.Buf)
		})
	}
}

func (s *KeeperMockVmTestSuite) TestDSServer_GetNativeBalance() {
	ctx := s.ctx

	accXfiBalance := sdk.NewInt(500)
	acc, _ := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, accXfiBalance))

	// fail: input
	{
		// address
		{
			s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
				resp, err := client.GetNativeBalance(ctx, &dvmTypes.NativeBalanceRequest{Address: []byte{0x0}, Ticker: "xfi"})
				s.Require().NoError(err)
				s.Require().NotNil(resp)

				s.Require().Equal(dvmTypes.ErrorCode_BAD_REQUEST, resp.ErrorCode)
				s.Require().NotEmpty(resp.ErrorMessage)
			})
		}

		// denom
		{
			s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
				resp, err := client.GetNativeBalance(ctx, &dvmTypes.NativeBalanceRequest{Address: types.Bech32ToLibra(acc.GetAddress()), Ticker: ""})
				s.Require().NoError(err)
				s.Require().NotNil(resp)

				s.Require().Equal(dvmTypes.ErrorCode_BAD_REQUEST, resp.ErrorCode)
				s.Require().NotEmpty(resp.ErrorMessage)
			})
		}
	}

	// ok: non-existing
	{
		accAddr, _, _ := tests.GenAccAddress()

		s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
			resp, err := client.GetNativeBalance(ctx, &dvmTypes.NativeBalanceRequest{Address: types.Bech32ToLibra(accAddr), Ticker: "mybtc"})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvmTypes.ErrorCode_NONE, resp.ErrorCode)
			s.Require().Empty(resp.ErrorMessage)
			s.checkVmU128Value(sdk.ZeroInt(), resp.Balance)
		})
	}

	// ok: existing acc: existing denom
	{
		s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
			resp, err := client.GetNativeBalance(ctx, &dvmTypes.NativeBalanceRequest{Address: types.Bech32ToLibra(acc.GetAddress()), Ticker: dnConfig.MainDenom})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvmTypes.ErrorCode_NONE, resp.ErrorCode)
			s.Require().Empty(resp.ErrorMessage)
			s.checkVmU128Value(accXfiBalance, resp.Balance)
		})
	}

	// ok: existing acc: non-existing denom
	{
		s.DoDSClientRequest(func(ctx context.Context, client dvmTypes.DSServiceClient) {
			resp, err := client.GetNativeBalance(ctx, &dvmTypes.NativeBalanceRequest{Address: types.Bech32ToLibra(acc.GetAddress()), Ticker: "mybtc"})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvmTypes.ErrorCode_NONE, resp.ErrorCode)
			s.Require().Empty(resp.ErrorMessage)
			s.checkVmU128Value(sdk.ZeroInt(), resp.Balance)
		})
	}
}

// That one is covered by integ tests
//func (s *KeeperMockVmTestSuite) TestDSServer_GetOraclePrice()

func (s *KeeperMockVmTestSuite) checkVmU128Value(expected sdk.Int, received *dvmTypes.U128) {
	s.Require().NotNil(received)
	s.Require().Len(received.Buf, 16)

	expectedU128, err := types.SdkIntToVmU128(expected)
	s.Require().NoError(err)
	s.Require().ElementsMatch(expectedU128.Buf, received.Buf)
}
