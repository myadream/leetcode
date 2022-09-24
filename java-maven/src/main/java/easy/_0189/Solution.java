package easy._0189;

import java.util.Arrays;

public class Solution {
    public void rotateArraySplice(int[] nums, int k) {
        int n = nums.length;
        int[] newArr = new int[n];
        for (int i = 0; i < n; ++i) {
            newArr[(i + k) % n] = nums[i];
        }
        System.arraycopy(newArr, 0, nums, 0, n);

    }
}
