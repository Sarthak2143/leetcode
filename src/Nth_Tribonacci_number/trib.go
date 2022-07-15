package main

func tribonacci(n int) int {
    t0, t1, t2 := 0, 1, 1
    for i := 0; i < n-2; i++ {
        temp1 := t0
        temp2 := t1
        t0 = t1
        t1 = t2
        t2 += temp1 + temp2
    }
    if n != 0{
        return t2
    } else{
        return 0
    }
}
