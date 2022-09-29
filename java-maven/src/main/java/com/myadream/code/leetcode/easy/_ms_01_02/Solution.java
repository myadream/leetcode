package com.myadream.code.leetcode.easy._ms_01_02;

import java.util.Arrays;
import java.util.HashMap;

public class Solution {
    public boolean CheckPermutationHashTableCount(String s1, String s2) {
        HashMap<Character, Integer> map = new HashMap<>(s1.length());
        for (int i = 0; i < s1.length(); i++) {
            map.put(s1.charAt(i), map.getOrDefault(s1.charAt(i), 0) + 1);
        }

        for (int i = 0; i < s2.length(); i++) {
            map.put(s2.charAt(i), map.getOrDefault(s2.charAt(i), 0) - 1);
            if (map.getOrDefault(s2.charAt(i), -1) == 0) {
                map.remove(s2.charAt(i));
            }
        }

        return map.isEmpty();
    }

    public boolean CheckPermutationHashTable(String s1, String s2) {
        if (s1.length() != s2.length()) {
            return false;
        }
         char[] c1 = new char[s1.length()];
         char[] c2 = new char[s2.length()];

         for (int i = 0; i < s1.length(); i++) {
             c1[i] = s1.charAt(i);
         }

         for (int i = 0; i < s2.length(); i++) {
             c2[i] = s2.charAt(i);
         }

        Arrays.sort(c1);
        Arrays.sort(c2);
        for (int i = 0; i < s1.length(); i++) {
            if (c1[i] != c2[i]) {
               return false;
            }
        }

        return true;
    }
}
