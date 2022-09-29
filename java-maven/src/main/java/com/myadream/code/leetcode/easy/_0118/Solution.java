package com.myadream.code.leetcode.easy._0118;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> list = new ArrayList<>(numRows);

        for (int i = 0; i < numRows; i++) {
            ArrayList<Integer> temp = new ArrayList<Integer>(i + 1);
            for (int j = 0; j < i + 1; j++) {
                if (j == 0 || i == j) {
                    temp.add(1);
                } else {
                    temp.add(list.get(i - 1).get(j - 1) + list.get(i - 1).get(j));
                }
            }

            list.add(temp);
        }

        return list;
    }
}
