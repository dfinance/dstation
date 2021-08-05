// fibonacci module
module Fibonacci {

   public fun calculate(n: u64): u64 {
                if (n == 1 || n == 2) {
                	return 1
                };

                if (n == 3) {
                    return 2
                };

                let a = calculate(n -1);
                let b = calculate(n - 2);

                a + b
            }

    }