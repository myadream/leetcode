## 方案一

常规解题思路，不断进行 `n / 2` 或者 `n / 2 + 1` 进行降低 `n` 的值

```go
package numberOfMatches

func numberOfMatches(n int) int {
	var res int

	for n > 1 {
		if n % 2 == 0 {
			n = n / 2
			res += n
		} else if n != 1 {
			n = ( n - 1 ) / 2

			res += n
			n += 1
		}
	}

	return res
}
```


## 方案二

所有比赛的和 `n - 1`,  所以直接返回 `n - 1` 的值
```go
package numberOfMatches

func numberOfMatches(n int) int {
	return n - 1
}

```

## 方案三

使用位运算 `n` 除以 2 `n >> 1`， 判断是否加 1 用按位计算 `n & 1`

```go
package numberOfMatches

func numberOfMatches(n int) int {
	var res int

	for n > 1 {
		res += n >> 1
		n = (n >> 1) + (n & 1)
	}

	return res
}
```