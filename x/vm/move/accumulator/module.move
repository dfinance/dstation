// Generic module with resource creation and modification.
module Accumulator {
    resource struct Res {
        sum: u64
    }

    public fun add(sender: &signer, value: u64): u64 acquires Res {
        let owner = 0x1::Signer::address_of(sender);

        if (!exists<Res>(owner)) {
            move_to<Res>(sender, Res {
                sum: 0
            })
        };

        let res = borrow_global_mut<Res>(owner);
        res.sum = res.sum + value;

        res.sum
    }
}