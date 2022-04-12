package easy._0001;

import java.util.Arrays;
import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> mark = new HashMap<Integer, Integer>();

        for (int i = 0; i < nums.length; i++) {
            if (mark.containsKey(nums[i])) {
                return new int[]{mark.get(nums[i]), i};
            }

            mark.put(target -nums[i] , i);
        }

        return null;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] nums = new int[]{2, 7, 11, 15};
        int target = 9;
        System.out.println(Arrays.toString(solution.twoSum(nums, target)));
    }
}