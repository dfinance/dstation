// Script excepts all arg types and asserts them.
script {
    use 0x1::Vector;

    fun main(account: &signer, arg_u8: u8, arg_u64: u64, arg_u128: u128, arg_addr: address, arg_bool_true: bool, arg_bool_false: bool, arg_vector: vector<u8>) {
        assert(arg_u8 == 128, 10);
        assert(arg_u64 == 1000000, 11);
        assert(arg_u128 == 100000000000000000000000000000, 12);

        assert(0x1::Signer::address_of(account) == arg_addr, 20);

        assert(arg_bool_true == true, 30);
        assert(arg_bool_false == false, 31);

        assert(Vector::length<u8>(&mut arg_vector) == 2, 40);
        assert(Vector::pop_back<u8>(&mut arg_vector) == 1, 41);
        assert(Vector::pop_back<u8>(&mut arg_vector) == 0, 42);
    }
}