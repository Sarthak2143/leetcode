def tribonacci(self, n: int) -> int:
    t0, t1, t2 = 0, 1, 1
    for _ in range(n-2):
        temp1 = t0
        temp2 = t1
        t0 = t1
        t1 = t2
        t2 = t2 + temp2 + temp1

    if n!=0:
        return t2
    else:
        return 0
