package com.myadream.code.leetcode.easy._0557;

public class Solution {
    public String reverseWordsString(String s) {
        String[] sp = s.split(" ");
        StringBuilder ans = new StringBuilder();

        for (int i = 0; i < sp.length; i++) {
            for (int j = sp[i].length() - 1; j >= 0; j--) {
                ans.append(sp[i].charAt(j));
            }

            ans.append(" ");
        }

        return ans.toString().trim();
    }
}
