//
// Created by Administrator on 2022/12/10 0010.
//

#include "maxSubArray.h"

int max(int a, int b) {
    if (a > b) return a;
    else return b;
}

int maxSubArray(int* nums, int numsSize) {
    int pre = 0;
    int maxAns = nums[0];

    for (int i = 0; i < numsSize; i++) {
        pre = fmax(pre + nums[i], nums[i]);
        maxAns = fmax(pre, maxAns);
    }

    return maxAns;
}


