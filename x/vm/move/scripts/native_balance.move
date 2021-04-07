script {
    use 0x1::Account;
    use 0x1::XFI;
    use 0x1::Event;

    fun main(account: &signer, sdk_balance: u128) {
        let move_balance = Account::get_native_balance<XFI::T>(account);

        Event::emit<u128>(account, move_balance);
        Event::emit<u128>(account, sdk_balance);
        assert(move_balance == sdk_balance, 1);
    }
}