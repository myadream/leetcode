package com.myadream.code.leetcode.medium._0870;

public class Solution {
    public int[] advantageCount(int[] nums1, int[] nums2) {
        int len = nums1.length;
        int[] ans = new int[len];

        for (int i = 0; i < len; i++) {
            int min = 0;
            for (int l = len; l > i; l--) {
                if (nums1[l] > nums2[i]) {
                   min = nums1[l];
                }

                
            }

            ans[i] = min;
        }

        return ans;
    }
}
