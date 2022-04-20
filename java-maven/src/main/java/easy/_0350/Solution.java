package easy._0350;

import java.lang.reflect.Array;
import java.util.*;

public class Solution {
    public int[] intersect(int[] nums1, int[] nums2) {
        int[] result = new int[nums2.length];
        ArrayList<Integer> list = new ArrayList<>(nums2.length / 2);

        for (int i = 0; i < nums1.length; i++) {
            list.add(nums1[i]);
        }

        int ri = 0;
        for (int i = 0; i < nums2.length; i++) {
            int index = list.indexOf(nums2[i]);
            if (index != -1) {
                list.remove(index);
                result[ri++] = nums2[i];
            }
        }

        return Arrays.copyOfRange(result, 0, ri);
    }
}
