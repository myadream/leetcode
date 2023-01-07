package com.myadream.code.leetcode.easy._1624;

import java.util.HashMap;

public class Solution {

    //搜索前后位置比较
    public int maxLengthBetweenEqualCharactersString(String s) {
        int maxLen = -1;

        for (int i = 97; i < 123; i++) {
            maxLen = Math.max(maxLen, s.lastIndexOf(i) - s.indexOf(i) - 1);
        }

        return maxLen;
    }

     //暴力破解
    public int maxLengthBetweenEqualCharactersHasTable(String s) {
        HashMap<Character, Integer> map = new HashMap<>(s.length() >> 1);
        int maxLen = -1;

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (map.containsKey(c)) {
               maxLen = Math.max(maxLen, i - map.get(c) - 1);
            } else {
                map.put(c, i);
            }
        }

        return maxLen;
    }

    //暴力破解
    public int maxLengthBetweenEqualCharactersViolent(String s) {
        int maxLen = -1;
        int len = s.length();

        for (int i = 0; i < len; i++) {
            for (int j = i + 1; j < len; j++) {
                if (s.charAt(i) == s.charAt(j)) {
                    maxLen = Math.max(maxLen, j - i - 1);
                }
            }
        }

        return maxLen;
    }
}
