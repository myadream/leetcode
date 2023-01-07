package com.myadream.code.leetcode.easy._0020;

public class Solution {
    public boolean isValid(String s) {
        int len = s.length();
        for (int i = 0; i < len; i ++) {
            if (s.isEmpty()) {
                return true;
            }
            s = s.replace("()", "").replace("{}", "").replace("[]", "");
        }

        return s.isEmpty();
    }
}