package com.myadream.code.leetcode.easy._0001;

import java.util.HashMap;
import java.util.Map;

class Solution {
    //暴力循环
    public int[] twoSum(int[] nums, int target) {
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[i] + nums[j] == target) {
                    return new int[]{i, j};
                }
            }
        }

        return new int[]{0, 0};
    }

    //hast Map
    public int[] twoSum2(int[] nums, int target) {
        Map<Integer, Integer> maps = new HashMap<>(nums.length - 1);

        for (int i = 0; i < nums.length; i++) {
            if (maps.containsKey(target - nums[i])) {
                return new int[]{maps.get(target - nums[i]), i};
            } else {
                maps.put(nums[i], i);
            }
        }

        return new int[]{0, 0};
    }


}
