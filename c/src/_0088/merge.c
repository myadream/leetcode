//
// Created by Administrator on 2022/12/13 0013.
//

#include "stdio.h"
#include "merge.h"
#include "../library/tools.h"

void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n) {
    for (int i = n; i < nums1Size; i++) {
        nums1[n] = nums2[i - n];
    }

    for (int i = 0; i < nums1Size; i++) {
        printf("nums[%d]: %d", i, nums1[i]);
    }

    qsort(nums1, nums1Size, sizeof(nums1[0]), sortByAsc);
}