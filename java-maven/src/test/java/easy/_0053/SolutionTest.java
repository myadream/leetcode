package easy._0053;

import easy.DataSetControl;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest  {

    Solution solution = new Solution();

    @Test
    void maxSubArray() {
        int[] nums = new int[]{-2,1,-3,4,-1,2,1,-5,4};
        int result = solution.maxSubArray(nums);

        assertEquals(6, result);

        nums = new int[]{1};
        result = solution.maxSubArray(nums);
        assertEquals(1, result);

        nums = new int[]{5,4,-1,7,8};
        result = solution.maxSubArray(nums);
        assertEquals(23, result);
    }
}