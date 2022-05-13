package easy._0217;

import java.util.Arrays;
import java.util.HashSet;

public class Solution {
    /**
     * 暴力循环
     *
     * @param nums
     * @return
     */
    public boolean containsDuplicate(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[i] == nums[j]) {
                    return true;
                }
            }
        }

        return false;
    }

    /**
     * 先排序后在比较
     *
     * @param nums
     * @return
     */
    public boolean containsDuplicate3(int[] nums) {
        Arrays.sort(nums);

        for (int i = 1; i < nums.length; i++) {
            if (nums[i - 1] == nums[i]) {
                return true;
            }
        }

        return false;
    }

    /**
     * hash table 处理
     *
     * @param nums
     * @return
     */
    public boolean containsDuplicate2(int[] nums) {
        HashSet<Integer> hashSet = new HashSet<>(nums.length / 2);

        for (int i = 0; i < nums.length; i++) {
            if (hashSet.contains(nums[i])) {
                return true;
            } else {
                hashSet.add(nums[i]);
            }
        }

        return false;
    }


    /**
     * 双指真
     *
     * @param nums
     * @return
     */
    public boolean containsDuplicate4(int[] nums) {
        int right = 0;
        int left = right + 1;

        while (nums.length > left) {
            if (nums[right] == nums[left]) {
                return true;
            } else {
                left++;
            }

            if (left >= nums.length) {
                right++;
                left = right + 1;
            }
        }

        return false;
    }

}
