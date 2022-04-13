package easy._0001;

import java.util.HashMap;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> mark = new HashMap<>(nums.length);

        for (int i = 0; i < nums.length; i++) {
            if (mark.containsKey(nums[i])) {
                return new int[]{mark.get(nums[i]), i};
            }

            mark.put(target - nums[i], i);
        }

        return new int[]{-1, -1};
    }

    /**
     * 暴力循环
     */
    public int[] twoSum2(int[] nums, int target) {
        int len = nums.length;

        for (int i = 0; i < len; i++) {
            for (int j = 0; j < len; j++) {
                if (nums[i] + nums[j] == target) {
                    return new int[]{i, j};
                }
            }
        }

        return new int[0];
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] nums = new int[]{2, 7, 11, 15};
        int target = 9;
        int[] result = new int[]{1, 1};

        assertArrayEquals(solution.twoSum(nums, target), result);
    }
}
