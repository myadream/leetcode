package easy._0122;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class SolutionTest {

    Solution solution = new Solution();

    @Test
    void maxProfit() {
        int[] nums = {7, 1, 5, 3, 6, 4};
        assertEquals(5, solution.maxProfit(nums));

        nums = new int[]{7, 6, 4, 3, 1};
        assertEquals(0, solution.maxProfit(nums));
    }
}