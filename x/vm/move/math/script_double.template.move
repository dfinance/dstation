// File contains multiple scripts (this is OK).
// Execution should be performed one by one OR withing single Tx (separate Messages).

script {
    use 0x1::Event;
    use %s::DblMathAdd;
    use %s::DblMathSub;

    fun main(account: &signer, a: u64, b: u64, c: u64) {
        let ab = DblMathAdd::add(a, b);
        let res = DblMathSub::sub(ab, c);
        Event::emit<u64>(account, res);
    }
}

script {
    use 0x1::Event;
    use %s::DblMathAdd;
    use %s::DblMathSub;

    fun main(account: &signer, a: u64, b: u64, c: u64) {
        let ab = DblMathSub::sub(a, b);
        let res = DblMathAdd::add(ab, c);
        Event::emit<u64>(account, res);
    }
}