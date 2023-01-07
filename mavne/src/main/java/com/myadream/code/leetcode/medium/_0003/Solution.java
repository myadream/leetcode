package com.myadream.code.leetcode.medium._0003;

public class Solution {
    public int lengthOfLongestSubstring(String s) {
        int ans = 0;
        int i = 0;
        int[] ascii = new int[127];

        for (int j = 0; j < s.length(); j++) {
            i = Math.max(i, ascii[s.charAt(j)]);
            ascii[s.charAt(j)] = j + 1;
            ans = Math.max(ans, j - i + 1);
        }

        return ans;
    }
}
