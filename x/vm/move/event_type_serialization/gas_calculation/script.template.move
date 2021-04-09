// Script triggers GasEvent module function.
script {
    use %s::GasEvent;

    fun main(account: &signer) {
        GasEvent::test(account);
    }
}