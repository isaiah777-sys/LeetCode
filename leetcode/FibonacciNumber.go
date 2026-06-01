package leetcode

func fib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    prev2 := 0
    prev1 := 1
    current := 0

    for i := 2; i <= n; i++ {
        current = prev1 + prev2
        prev2 = prev1
        prev1 = current
    }

    return current
}