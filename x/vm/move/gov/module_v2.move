// Generic StdLib Move module with a fix.
address 0x1 {
    module StdMath {
        public fun inc(v: u64): u64 {
            v + 1
        }
    }
}