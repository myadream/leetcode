#include "removeDuplicates.h"
#include "stdio.h"

int removeDuplicates(int* nums, int numsSize) {
    int res = 1, i = 1;

    if (numsSize == 1) {
        return res;
    }

    for (;i < numsSize; i++) {
        if (nums[i] == nums[i - 1]) {
            continue;
        }

        nums[res] = nums[i];
        res++;
    }

    return res;
}