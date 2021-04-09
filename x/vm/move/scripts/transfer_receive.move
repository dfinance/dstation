// Scripts withdraws signer coins from VM to SDK.
script {
    use 0x1::Account;
    use 0x1::XFI;
    use 0x1::Coins;

    fun main(account: &signer, xfi_amount: u128, eth_amount: u128, btc_amount: u128, usdt_amount: u128) {
        // Acquire tokens from resource
        let xfi_token = Account::withdraw_from_sender<XFI::T>(account, xfi_amount);
        let eth_token = Account::withdraw_from_sender<Coins::ETH>(account, eth_amount);
        let btc_token = Account::withdraw_from_sender<Coins::BTC>(account, btc_amount);
        let usdt_token = Account::withdraw_from_sender<Coins::USDT>(account, usdt_amount);

        // Witdraw signer coins
        Account::withdraw_native<XFI::T>(account, xfi_token);
        Account::withdraw_native<Coins::ETH>(account, eth_token);
        Account::withdraw_native<Coins::BTC>(account, btc_token);
        Account::withdraw_native<Coins::USDT>(account, usdt_token);
    }
}