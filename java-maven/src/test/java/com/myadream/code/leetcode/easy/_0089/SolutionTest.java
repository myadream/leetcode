package com.myadream.code.leetcode.easy._0089;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

class SolutionTest {
    Solution solution = new Solution();

    @Test
    void merge() {
        int[] nums1 = new int[]{1,2,3,0,0,0};
        int[] nums2 = new int[]{2,5,6};
        int m = 3;
        int n = 3;

        solution.merge(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1,2,2,3,5,6}, nums1);

        nums1 = new int[]{1};
        nums2 = new int[]{};
        m = 1;
        n = 0;

        solution.merge(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1}, nums1);


        nums1 = new int[]{0};
        nums2 = new int[]{1};
        m = 0;
        n = 1;

        solution.merge(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1}, nums1);

        nums1 = new int[]{1,2,4,5,6,0};
        nums2 = new int[]{3};
        m = 5;
        n = 1;

        solution.merge(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1,2,3,4,5,6}, nums1);

        nums1 = new int[]{0,0,0,0,0};
        nums2 = new int[]{1,2,3,4,5};
        m = -0;
        n = 5;

        solution.merge(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1,2,3,4,5}, nums1);
    }

    @Test
    void merge2() {
        int[] nums1 = new int[]{1,2,3,0,0,0};
        int[] nums2 = new int[]{2,5,6};
        int m = 3;
        int n = 3;

        solution.merge2(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1,2,2,3,5,6}, nums1);

        nums1 = new int[]{1};
        nums2 = new int[]{};
        m = 1;
        n = 0;

        solution.merge2(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1}, nums1);


        nums1 = new int[]{0};
        nums2 = new int[]{1};
        m = 0;
        n = 1;

        solution.merge2(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1}, nums1);

        nums1 = new int[]{1,2,4,5,6,0};
        nums2 = new int[]{3};
        m = 5;
        n = 1;

        solution.merge2(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1,2,3,4,5,6}, nums1);

        nums1 = new int[]{0,0,0,0,0};
        nums2 = new int[]{1,2,3,4,5};
        m = -0;
        n = 5;

        solution.merge2(nums1, m, nums2, n);
        assertArrayEquals(new int[]{1,2,3,4,5}, nums1);
    }
}