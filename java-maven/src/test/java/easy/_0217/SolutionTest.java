package easy._0217;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    Solution solution = new Solution();

    @Test
    void containsDuplicate() {
        int[] nums = new int[] {1,2,3,1};
        boolean result = solution.containsDuplicate(nums);

        assertTrue(result);

        nums = new int[]{1,2,3,4};
        result = solution.containsDuplicate(nums);

        assertFalse(result);


        nums = new int[]{1,1,1,3,3,4,3,2,4,2};
        result = solution.containsDuplicate(nums);

        assertTrue(result);
    }

    @Test
    void containsDuplicate1() {
        int[] nums = new int[] {1,2,3,1};
        boolean result = solution.containsDuplicate1(nums);

        assertTrue(result);

        nums = new int[]{1,2,3,4};
        result = solution.containsDuplicate1(nums);

        assertFalse(result);


        nums = new int[]{1,1,1,3,3,4,3,2,4,2};
        result = solution.containsDuplicate1(nums);

        assertTrue(result);
    }

    @Test
    void containsDuplicate2() {
        int[] nums = new int[] {1,2,3,1};
        boolean result = solution.containsDuplicate2(nums);

        assertTrue(result);

        nums = new int[]{1,2,3,4};
        result = solution.containsDuplicate2(nums);

        assertFalse(result);


        nums = new int[]{1,1,1,3,3,4,3,2,4,2};
        result = solution.containsDuplicate2(nums);

        assertTrue(result);
    }
}