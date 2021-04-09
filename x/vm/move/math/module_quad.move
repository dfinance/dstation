// File contains multiple Move modules (this is OK).

module DblMathAdd {
    public fun add(a: u64, b: u64): u64 {
        a + b
    }
}

module DblMathSub {
    public fun sub(a: u64, b: u64): u64 {
        a - b
    }
}

module DblMathMul {
    public fun sub(a: u64, b: u64): u64 {
        a * b
    }
}

module DblMathQuo {
    public fun sub(a: u64, b: u64): u64 {
        a / b
    }
}