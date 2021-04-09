// Module emits complex events with 4-deep nested struct.
module GasEvent {
    struct A {
        value: u64
    }

    struct B<T> {
        value: T
    }

    struct C<T> {
        value: T
    }

    struct D<T> {
        value: T
    }

    public fun test(account: &signer) {
        let a = A {
            value: 10
        };

        let b = B<A> {
            value: a
        };

        let c = C<B<A>> {
            value: b
        };

        let d = D<C<B<A>>> {
            value: c
        };

        0x1::Event::emit<D<C<B<A>>>>(account, d);
    }

}