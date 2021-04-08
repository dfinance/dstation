// Script triggers StdMath module inc function.
script {
    use 0x1::StdMath;

    fun main(account: &signer, v: u64) {
        let res = StdMath::inc(v);

        0x1::Event::emit<u64>(account, res);
    }
}