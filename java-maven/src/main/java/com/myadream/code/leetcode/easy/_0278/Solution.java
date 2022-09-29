package com.myadream.code.leetcode.easy._0278;

public class Solution {

    private int bad;

    public void setBad(int bad) {
        this.bad = bad;
    }

    public boolean isBadVersion(int bad) {
        return this.bad == bad;
    }

    public int firstBadVersion(int n) {
        int left = 1;
        int right = n;

        while (left < right)  {
            int mid = left + (right - left) >> 2;
            if (isBadVersion(mid)) {
               right = mid;
            } else {
                left = mid + 1;
            }
        }

        return left;
    }
}
