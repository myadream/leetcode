//
// Created by gaojian on 2022/8/12.
//

#include "firstBadVersion.h"
//
//int firstBadVersion(int n) {
//    int left = 0;
//    int mid = 0;
//
//    while (left < n) {
//        mid = left + (n - left) / 2 ;
//       if (isBadVersion(mid)) {
//           n = mid;
//       } else {
//           left = mid + 1;
//       }
//    }
//
//    return left;
//}

int firstBadVersion(int n) {
    long long first = 1, last = (long long ) n + 1;
    long long mid = first + (last - first) / 2;
    while (isBadVersion(mid) && !isBadVersion(mid - 1)) {
        mid = first + (last - first) / 2;
        if (!isBadVersion(mid)) {
            first = mid;
        } else {
            last = mid;
        }
    }

    return 0;
}

int binarySearch(long long left, long long right) {
    long long mid = left + (left - right) / 2;

    if (isBadVersion(mid) && isBadVersion(mid - 1)) {
      return mid;
    }

    if (!isBadVersion(mid)) {
        return binarySearch(mid + 1, right);
    }

    return binarySearch(left, mid);
}