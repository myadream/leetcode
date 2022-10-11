//
// Created by Administrator on 2022/10/11 0011.
//

#include <malloc.h>
#include "sortedSquares.h"


int *sortedSquaresTwoPointers(int *nums, int numsSize, int *returnSize) {
    int* ans = malloc(sizeof(int) * numsSize);
    *returnSize = numsSize;

    for (int i = 0, j = numsSize - 1, pos = numsSize - 1; i <= j;) {
        if (square(nums[i]) > square(nums[j])) {
            ans[pos] = square(nums[i++]);
        } else {
            ans[pos] = square(nums[j--]);
        }
        --pos;
    }

    return ans;
}

int square(int x) {
    return x < 0 ? -(x * x) : x * x;
}

