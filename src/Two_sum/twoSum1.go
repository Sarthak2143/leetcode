package main

func twoSum(nums []int, target int) []int{
    var re []int
    for i, j := range nums{
        for x, y := range nums{
            if i == x{
                continue
            } else if y + j == target{
                re = []int(x, i)
                break
            } else{
                continue
            }
        }
    }
    return re
}
