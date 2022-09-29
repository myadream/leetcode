package medium._0167;

public class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int l = 0;
        int r = numbers.length - 1;
        int ans = target - numbers[l];

        while (l < r) {
            if (ans == numbers[r]) {
                return new int[]{l + 1, r + 1};
            } else if (numbers[r] - ans < 0) {
                l++;
                ans = target - numbers[l];
            } else {
                r--;
            }
        }

        return new int[]{};
    }
}
