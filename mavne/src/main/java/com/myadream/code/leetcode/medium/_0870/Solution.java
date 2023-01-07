package com.myadream.code.leetcode.medium._0870;

import java.lang.reflect.Array;
import java.util.Arrays;

public class Solution {
    public int[] advantageCount(int[] nums1, int[] nums2) {
        int len = nums1.length;
        int[][] t = new int[len][2];

        for (int i = 0; i < len; i++) {
            t[i] = new int[]{nums2[i], i};
        }

        Arrays.sort(nums1);
        Arrays.sort(t, (a, b) -> a[0] - b[0]);

        int[] ans = new int[len];
        int i = 0, j = len - 1;

        for (int v: nums1) {
            if (v <= t[i][0]) {
                ans[t[j--][1]] = v;
            } else {
                ans[t[i++][1]] = v;
            }
        }

        return ans;
    }
}
