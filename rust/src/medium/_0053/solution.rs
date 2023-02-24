struct Solution;

impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut ans = nums[0];
        let mut sum = nums[0];
        nums.iter().skip(1).for_each(|&x| {
            sum = x.max(sum + x);
            ans = ans.max(sum);
        });
        ans
    }
}

#[cfg(test)]
mod test {
    use std::vec;

    use super::Solution;

    #[test]
    fn test() {
        assert_eq!(Solution::max_sub_array(vec![-2,1,-3,4,-1,2,1,-5,4]), 6);
        assert_eq!(Solution::max_sub_array(vec![1]), 1);
        assert_eq!(Solution::max_sub_array(vec![5,4,-1,7,8]), 23);
    }
}