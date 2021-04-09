// Module creates FooEvent which emits some event.
module Foo {
    struct FooEvent<T, VT> {
        field_T:  T,
        field_VT: VT
    }

    public fun NewFooEvent<T, VT>(account: &signer, arg_T: T, arg_VT: VT): FooEvent<T, VT> {
        let fooEvent = FooEvent<T, VT> {
            field_T:  arg_T,
            field_VT: arg_VT
        };

        0x1::Event::emit<bool>(account, true);

        fooEvent
    }
}