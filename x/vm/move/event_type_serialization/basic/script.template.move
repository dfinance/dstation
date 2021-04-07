// Script emits various events (including the one from the Foo module.
script {
    use %s::Foo;

    fun main(account: &signer) {
        // Event with single tag
        0x1::Event::emit<u8>(account, 128);

        // Event with single vector
        0x1::Event::emit<vector<u8>>(account, x"0102");

        // Two events:
        //   1. Module: single tag
        //   2. Script: generic struct with tag, vector
        let fooEvent = Foo::NewFooEvent<u64, vector<u8>>(account, 1000, x"0102");
        0x1::Event::emit<Foo::FooEvent<u64, vector<u8>>>(account, fooEvent);
    }
}