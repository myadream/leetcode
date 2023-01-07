package com.myadream.code.leetcode.easy._0283;


public class Solution {
    public void moveZeroesTwoPointers(int[] nums) {
        int l = 0;
        int r = 1;

        while (r < nums.length) {
            if (nums[r] != 0) {
                int t = nums[l];
                nums[l] = nums[r];
                nums[r] = t;
                l++;
            }
            r++;
        }
    }


    public void moveZeroesArray(int[] nums) {
        int j = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != 0) {
                nums[j++] = nums[i];
            }
        }

        for (int i = j; i < nums.length; i++) {
            nums[i] = 0;
        }

    }
}
