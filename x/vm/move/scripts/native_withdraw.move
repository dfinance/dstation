// Script withdraws balance from the VM resource and performs the VM->SDK transfer.
script {
    use 0x1::Account;
    use 0x1::Dfinance;
    use 0x1::XFI;

    fun main(account: &signer, amount: u128) {
        let token = Account::withdraw_from_sender<XFI::T>(account, amount);
        assert(Dfinance::value<XFI::T>(&token) == amount, 1);

        Account::withdraw_native<XFI::T>(account, token);
    }
}