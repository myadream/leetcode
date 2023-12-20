pub fn case_one(nums: Vec<i32>) -> i32 {
    let mut ans = nums[0];
    let mut sum = nums[0];
    nums.iter().skip(1).for_each(|&x| {
        sum = x.max(sum + x);
        ans = ans.max(sum);
    });
    ans
}
