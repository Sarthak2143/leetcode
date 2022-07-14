## Given:

[Click here](https://leetcode.com/problems/two-sum/) for official problem page. 

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `*target*`.

You may assume that each input would have *exactly one solution*, and you may not use the same element twice.

You can return the answer in any order.


Example 1:
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

Example 2:
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

Example 3:
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

### Constraints:

- `2 <= nums.length <= 104`
- `-10^9 <= nums[i] <= 109`
- `-10^9 <= target <= 109`
- **Only one valid answer exists.**


**Follow-up**: Can you come up with an algorithm that is less than `O(n^2)` time complexity?

__

## Solution

### Approach 1: Brute Force

**Algorithm**

The brute force approach is simple. Loop through each element `x` and find if there is another value that equals to `target - x`.

**Complexity Analysis**

- *Time complexity*: `O(n^2)`. For each element, we try to find its complement by looping through the rest of the array which takes `O(n)` time. Therefore, the time complexity is `O(n^2)`.

- *Space complexity*: `O(1)`. The space required does not depend on the size of the input array, so only constant space is used.

**Implementation**:

```go
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
```


