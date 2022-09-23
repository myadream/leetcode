package easy._0977;

import java.util.Arrays;

public class Solution {
    public int[] sortedSquaresSorting(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            nums[i] = squares(nums[i]);
        }

        Arrays.sort(nums);
        return nums;
    }

    public int[] sortedSquareTwoPointers(int[] nums) {
        int n = nums.length - 1;
        int[] ans = new int[nums.length];
        for (int i = 0, j = n, pos = n; i <= j;) {
            if (squares(nums[i]) > squares(nums[j])) {
                ans[pos] = squares(nums[i]);
                i++;
            } else {
                ans[pos] = squares(nums[j]);
                j--;
            }

            pos--;
        }

        return ans;
    }



    private int squares(int n) {
        return Math.abs(n * n);
    }
}
