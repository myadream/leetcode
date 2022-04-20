package easy._0053;

import java.util.Arrays;
import java.util.HashSet;

public class Solution {
    public int maxSubArray(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            nums[i] = nums[i] + Math.max(nums[i - 1], 0);
            nums[0] = Math.max(nums[i], nums[0]);
        }

        return nums[0];
    }
}
