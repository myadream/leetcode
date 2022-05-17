package easy._0350;

import java.util.ArrayList;
import java.util.Arrays;

public class Solution {
    //暴力
    public int[] intersect(int[] nums1, int[] nums2) {
        int[] any = new int[nums2.length];
        ArrayList<Integer> list =new ArrayList<>(nums2.length / 2);

        for (int i = 0; i < nums1.length; i++) {
            list.add(nums1[i]);
        }

        int flag = 0;
        for (int i = 0; i < nums2.length; i++) {
            int index = list.indexOf(nums2[i]);
            if (index != -1) {
               list.remove(index);
               any[flag++] = nums2[i];
            }
        }

        return Arrays.copyOfRange(any, 0, flag);
    }

    //双指针
    public int[] intersect2(int[] nums1, int[] nums2) {
        Arrays.sort(nums1);
        Arrays.sort(nums2);
        int[] any = new int[nums2.length];

        int left = 0;
        int right = 0;
        int flag = 0;

        while (left < nums1.length && right < nums2.length) {
           if (nums1[left] < nums2[right]) {
              left++;
           } else if (nums1[left] > nums2[right]) {
               right++;
           } else {
               any[flag++] = nums1[left];
               left++;
               right++;
           }
        }

        return Arrays.copyOfRange(any, 0, flag);
    }


}
