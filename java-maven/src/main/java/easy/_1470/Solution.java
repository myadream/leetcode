package easy._1470;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public int[] shuffle(int[] nums, int n) {
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            ans.add(nums[i]);
            ans.add(nums[i + n]);
        }

    }
}
