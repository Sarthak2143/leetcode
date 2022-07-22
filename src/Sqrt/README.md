## Given

[Click here](https://leetcode.com/problems/sqrtx/) for official problem page.

Given a *non-negative integer* `x`, compute and return *the square root of* `x`.

Since the return type is an integer, the decimal digits are **truncated**, and only **the integer part** of the result is returned.

**Note**: You are not allowed to use any built-in exponent function or operator, such as `pow(x, 0.5)` or `x ** 0.5`.


Example 1:
```
Input: x = 4
Output: 2
```

Example 2:
```
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.
```

## Constraints:

- `0 <= x <= 2^31 - 1`

### Solution

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

`z -= (z*z - x)/(2*x)`

Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

**Note**: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called [Newton's method](https://en.wikipedia.org/wiki/Newton%27s_method). It works well for many functions but especially well for square root.

**Golang Solution**

```go
func Sqrt(x int) int {
    z := 1.0

    for i := 0; i < 1000; i++{
        z -= (z*z - float64(x))/(2*z)
    }
    return int(z)
}
```

**Python solution**

```py
def sqrt(x):
    z = 1.0
    for i in range(1, 100):
        z -= (z*z - float(x))/(2*z)
    return int(z)
```

---
