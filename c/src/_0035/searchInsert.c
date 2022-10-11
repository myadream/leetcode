//
// Created by Administrator on 2022/10/10.
//

#include "searchInsert.h"

int searchInsert(int* nums, int numsSize, int target) {
    int l = 0;

    while (l < numsSize) {
        int mid  = l + (numsSize - l) / 2;
        if (nums[mid] > target) {
            numsSize = mid;
        } else {
            l = mid + 1;
        }
    }

    return l;
}