# This is a better solution
# Just use `n<=1` and return n

def fib(n):
    if n <= 1:
        return n
    return fib(n-1) + (n-2)
