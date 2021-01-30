#include "removeDuplicates.h"

#include "stdio.h"
#include "stdlib.h"

int main(void) {
    int nums[10] = {0,0,1,1, 1,2,2,3 ,3,4};
    int target = 5;

    int result = removeDuplicates(nums, 10);
    if (target != result) {
        printf("err: %d \n", result);
        return 0;
    }

    int nums1 = {1, 1, 2};
    target = 2;

    result = removeDuplicates(nums1, 3);
    if (target != result) {
        printf("err: %d \n", result);
        return 0;
    }
}