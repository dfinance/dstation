// Generic module with resources.
address 0x1 {
    module Foo {
        resource struct U64 {val: u64}
        resource struct Obj {val: u64, o: U64}

        public fun add(sender: &signer, a: u64, b: u64): u64 {
            let sum = a + b;
            let value = U64 {val: sum};
            move_to<U64>(sender, value);
            sum
        }

        public fun build_obj(sender: &signer, a: u64) {
            let u64val = U64 {val: a};
            let value = Obj {val: a, o: u64val};
            move_to<Obj>(sender, value);
        }
    }
}