module MixedModule {
    public fun add(a: u64, b: u64): u64 {
        a + b
    }
}

script {
    use 0x1::Event;

    fun main(account: &signer, a: u64, b: u64) {
        let c = a + b;
        Event::emit<u64>(account, c);
    }
}