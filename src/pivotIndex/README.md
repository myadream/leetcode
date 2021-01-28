## 解题思路

### 方案一

参考官方解题思路, 先统计所有数的和, 在第二次循环中从左往右进行比较

```go
	var total, sum int
	var n = len(nums)

	for i := 0; i < n; i++ {
		total += nums[i]
	}

	for i := 0; i < n; i++ {
		fmt.Println(2*sum+nums[i], total)
		if 2*sum+nums[i] == total {
			return i
		}

		sum += nums[i]
	}
```