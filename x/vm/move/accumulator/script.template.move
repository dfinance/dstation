// Script performs the Accumulator module call and emits event with the latest resource value.
script {
    use 0x1::Event;
    use %s::Accumulator;

    fun main(sender: &signer, value: u64) {
        let sum = Accumulator::add(sender, value);
        Event::emit<u64>(sender, sum);
    }
}