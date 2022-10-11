package com.myadream.code.leetcode.easy._1790;

public class Solution {
    public boolean areAlmostEqualCount(String s1, String s2) {
        int count = 0;
        char swap = 0;
        char swap2 = 0;
        for (int i = 0; i < s1.length(); i++) {
            if (s1.charAt(i) != s2.charAt(i)) {
                count++;

                if (swap != 0 && (s2.charAt(i) != swap || s1.charAt(i) != swap2)) {
                    return false;
                }

                if (swap == 0) {
                    swap2 = s2.charAt(i);
                    swap = s1.charAt(i);
                }
            }
        }

        return count == 2 || count == 0;
    }
}
