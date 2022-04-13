package easy._0001;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {
    Solution solution = new Solution();

    @Test
    void twoSum() {
        int[] nums = new int[]{2, 7, 11, 15};
        int target = 9;
        int[] result = new int[]{0, 1};

        assertEquals(Arrays.toString(solution.twoSum(nums, target)), Arrays.toString(result));
    }

    @Test
    void twoSum2() {
    }
}