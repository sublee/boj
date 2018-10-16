function fib(n) {
    if (n <= 1) {
        return n
    }
    if (n in m) {
        return m[n]
    }
    m[n] = fib(n-1) + fib(n-2)
    return m[n]
}

{
    print fib($1)
}
