// Module emits complex events with 6-deep nested struct and should fail with "out-of-gas".
module OutOfGasEvent {
    struct A {
        value: u64
    }

    struct B<T> {
        value: T
    }

    struct C<T> {
        value: T
    }

    struct Z<T> {
        value: T
    }

    struct V<T> {
        value: T
    }

    struct M<T> {
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

        let z = Z<C<B<A>>> {
            value: c
        };

        let v = V<Z<C<B<A>>>> {
            value: z
        };

        let m = M<V<Z<C<B<A>>>>> {
            value: v
        };

        0x1::Event::emit<M<V<Z<C<B<A>>>>>>(account, m);
    }

}