## 解题思路

### 方案一

使用 `strconv.Itoa`  函数转换为字符串统计字符长度。 在判断是否2的倍数

```go
if len(strconv.Itoa(nums[i])) % 2 == 0 {
    res++
}
```

### 方案二

循环 `nums` 变量在不断除以 10 移动整数位置每除以一次数量加 1 判断总共除以数量
```go
t := 0

for nums[i] != 0 {
    nums[i] /= 10
    t++
}

if t % 2 == 0 {
    res++
}
```