// Script performs a simple assert.
script {
    fun main(_account: &signer, value: u64) {
        assert(value == 1000, 1);
    }
}