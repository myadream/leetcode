package com.myadream.code.leetcode.medium._0713;

//todo 待学习
public class Solution {
    public int numSubarrayProductLessThanK(int[] nums, int k) {

        int ans = 0, left = 0, cur = 1;
        for (int right = 0; right < nums.length; right++) {
            cur *= nums[right];
            while (left <= right && cur >= k)
                cur /= nums[left++];
            ans += right - left + 1;
        }
        return ans;
    }
}
