// Script requests multiple CurrencyInfo resources and asserts their decimals.
script {
    use 0x1::Dfinance;
    use 0x1::XFI;
    use 0x1::Coins;

    fun main(_account: &signer, xfi_decimals: u8, eth_decimals: u8, btc_decimals: u8, usdt_decimals: u8) {
        assert(Dfinance::decimals<XFI::T>() == xfi_decimals, 1);
        assert(Dfinance::decimals<Coins::ETH>() == eth_decimals, 2);
        assert(Dfinance::decimals<Coins::BTC>() == btc_decimals, 3);
        assert(Dfinance::decimals<Coins::USDT>() == usdt_decimals, 4);
    }
}