package com.myadream.code.leetcode.easy._0089;

import java.util.Arrays;

/**
 * @author gaojian
 */
public class Solution {

    /**
     * 常规循环
     *
     */
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int j = 0;

        while (j < n) {
            nums1[m + j] = nums2[j++];
        }

        Arrays.sort(nums1);
    }

    /**
     * 正向双指针
     *
     */
    public void merge2(int[] nums1, int m, int[] nums2, int n) {
        int p1 = 0;
        int p2 = 0;
        int current;
        int[] sorted = nums1.clone();

        while (p1 < m || p2 < n) {
            if (p1 == m) {
                current = nums2[p2++];
            } else if (p2 == n) {
                current = sorted[p1++];
            } else if (sorted[p1] < nums2[p2]) {
                current = sorted[p1++];
            } else {
                current = nums2[p2++];
            }

            nums1[p1 + p2 - 1] = current;
        }
    }
}
