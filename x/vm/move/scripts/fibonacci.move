script {
    use wallet1q924v95jw390p8ktjqzgh9s63q8pdm0aqpx3ey::Fibonacci;
    use 0x1::Event;

    fun main(account: &signer, n :u64) {
       let result = Fibonacci::calculate(n);

       Event::emit<u64>(account, result);
    }
}