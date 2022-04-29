package easy._0905;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public int[] sortArrayByParity(int[] nums) {
        List<Integer> ant = new ArrayList<>(nums.length);

        for (int num : nums) {
            if (num % 2 == 0) {
                ant.add(num);
            }
        }
        for (int num : nums) {
            if (num % 2 != 0) {
                ant.add(num);
            }
        }

        int[] m = new int[nums.length];
        for (int i = 0; i < ant.size(); i++) {
            m[i] = ant.get(i);
        }

        return m;
    }

    //常规双指真
    public int[] sortArrayByParity_2(int[] nums) {
        int left = 0;
        int right = nums.length - 1;
        int[] ant = new int[nums.length];

        for (int num : nums) {
            if (num % 2 == 0) {
                ant[left] = num;
                left++;
            } else {
                ant[right] = num;
                right--;
            }
        }

        return ant;
    }

    //双指针-优化
    public int[] sortArrayByParity_2_1(int[] nums) {
        //todo 待继续补充， 通过 i 当前位置进行计算得到存储位置
        int[] ant = new int[nums.length];
        for (int i = nums.length - 1; i >= 0; i--) {
            if (nums[i] % 2 == 0) {
                ant[Math.abs(i - (nums.length - 1))] = nums[i];
            } else {
                ant[i - (nums.length)] = nums[i];
            }
        }

        return ant;
    }

    public int[] sortArrayByParity_3(int[] nums) {
        int left = 0;
        int right = nums.length - 1;
        int t;

        while (left < right) {
            if (nums[left] % 2 != 0) {
                if (nums[right] % 2 == 0) {
                     t = nums[left];
                     nums[left] = nums[right];
                     nums[right] = t;
                } else {
                    right--;
                }
            } else {
                left++;
            }
        }

        return nums;
    }
}
