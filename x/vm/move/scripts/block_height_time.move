// Script receives current block time and height.
script {
    use 0x1::Block;
    use 0x1::Time;

    fun main(_account: &signer, sdk_block: u64, sdk_time: u64) {
        let native_block = Block::get_current_block_height();
        let native_time = Time::now();

        assert(native_block == sdk_block, 1);
        assert(native_time == sdk_time, 2);
    }
}