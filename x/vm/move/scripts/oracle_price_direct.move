script {
    use 0x1::Coins;

    fun main(_account: &signer, sdk_price: u128) {
        let native_price = Coins::get_price<Coins::ETH, Coins::USDT>();
        assert(native_price == sdk_price, 1);
    }
}