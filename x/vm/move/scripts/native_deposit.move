// Script performs the SDK->VM coin transfer and deposits it the VM resource.
script {
    use 0x1::Account;
    use 0x1::Dfinance;
    use 0x1::XFI;

    fun main(account: &signer, amount: u128) {
        let token = Account::deposit_native<XFI::T>(account, amount);
        assert(Dfinance::value<XFI::T>(&token) == amount, 1);

        Account::deposit_to_sender(account, token);
    }
}