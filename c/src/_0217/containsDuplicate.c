//
// Created by Administrator on 2022/12/10 0010.
//

#include "containsDuplicate.h"

int cmp(const void *a, const void *b) {
    return *(int *) a - * (int *) b;
}

// 排序比较
bool containsDuplicateQsort(int* nums, int numsSize) {
    qsort(nums, numsSize, sizeof (nums[0]), cmp);

    for (int i = 1; i < numsSize; i++) {
        if (nums[i] == nums[i - 1]) {
            return true;
        }
    }

    return false;
}
