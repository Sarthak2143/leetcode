func lenLongestFibSubseq(arr []int) int {
    set := make(map[int]bool)
    for _, v := range arr {
        set[v] = true
    }
    res := 0
    for i := 0; i < len(arr); i++{
        for j := i+1; j < len(arr); j++{
            a := arr[i]
            b := arr[j]
            c := a + b
            l := 2
            for set[c] {
                a = b
                b = c
                c = a + b
                l++
                res = max(res, l)

            }
        }
    }
    return res
}

func max(a, b int) int{
    if a > b { return a }
    return b
}
        
    

