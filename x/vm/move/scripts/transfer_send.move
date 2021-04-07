// Scripts deposits signer coins from SDK to VM and send them to the recepient.
script {
    use 0x1::Account;
    use 0x1::XFI;
    use 0x1::Coins;

    fun main(account: &signer, recipient: address, xfi_amount: u128, eth_amount: u128, btc_amount: u128, usdt_amount: u128) {
        // Deposit signer coins
        let xfi_token = Account::deposit_native<XFI::T>(account, xfi_amount);
        Account::deposit_to_sender(account, xfi_token);

        let eth_token = Account::deposit_native<Coins::ETH>(account, eth_amount);
        Account::deposit_to_sender(account, eth_token);

        let btc_token = Account::deposit_native<Coins::BTC>(account, btc_amount);
        Account::deposit_to_sender(account, btc_token);

        let usdt_token = Account::deposit_native<Coins::USDT>(account, usdt_amount);
        Account::deposit_to_sender(account, usdt_token);

        // Transfer to recipient resource
        Account::pay_from_sender<XFI::T>(account, recipient, xfi_amount);
        Account::pay_from_sender<Coins::ETH>(account, recipient, eth_amount);
        Account::pay_from_sender<Coins::BTC>(account, recipient, btc_amount);
        Account::pay_from_sender<Coins::USDT>(account, recipient, usdt_amount);
    }
}