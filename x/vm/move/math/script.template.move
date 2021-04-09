script {
    use 0x1::Event;
    use %s::Math;

    fun main(account: &signer, a: u64, b: u64) {
        let c = Math::add(a, b);
        Event::emit<u64>(account, c);
    }
}