#include "solution.h"

int search(int* nums, int numsSize, int target){
    int l = 0, r = numsSize;
    while (l < r) {
        int mid = l + (r - l) / 2;
        if (nums[mid] == target) {
            return mid;
        } else if (nums[mid] > target) {
            r--;
        } else {
            l++;
        }
    }

    return -1;
}