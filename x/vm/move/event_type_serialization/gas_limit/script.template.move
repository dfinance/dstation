// Script triggers OutOfGasEvent module function.
script {
    use %s::OutOfGasEvent;

    fun main(account: &signer) {
        OutOfGasEvent::test(account);
    }
}