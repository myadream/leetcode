package easy._0217;

import java.lang.reflect.Array;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;

public class Solution {
    public boolean containsDuplicate(int[] nums) {
        Arrays.sort(nums);
        int len = nums.length - 1;
        for (int i = 0; i < len; i++) {
            if (nums[i] == nums[i + 1]) {
                return true;
            }
        }

        return false;
    }

    /**
     * 使用map作为统计
     *
     * @param nums
     * @return
     */
    public boolean containsDuplicate1(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap<>(nums.length);
        for (int num:nums) {
            if (map.containsKey(num)) {
                return true;
            } else {
                map.put(num, 1);
            }
        }

        return false;
    }

    /**
     * 使用set作为统计
     *
     * @param nums
     * @return
     */
    public boolean containsDuplicate2(int[] nums) {
        HashSet<Integer> map = new HashSet<>(nums.length / 2);
        for (int num:nums) {
            if (!map.add(num)) {
               return true;
            }
        }

        return false;
    }
}
