package com.myadream.code.leetcode.easy._0455;

import java.util.Arrays;

public class Solution {
    public int findContentChildren(int[] g, int[] s) {
        Arrays.sort(g);
        Arrays.sort(s);

        int cookie = 0;
        int child = 0;

        while (child < g.length && cookie < s.length) {
            if (g[child] <= s[cookie]) {
                child++;
            }

            cookie++;
        }

        return child;
    }
}
