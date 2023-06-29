package leetcode278

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func defaultSearch(n int, bad int) int {
	for i := 1; i <= n; i++ {
		if isBadVersion(i, bad) {
			return i
		}
	}

	return 0
}

func binarySearch(n int, bad int) int {
	l := 0
	for l < n {
		mid := l + (n-l)/2
		if isBadVersion(mid, bad) {
			n = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func isBadVersion(n int, bad int) bool {
	return n == bad
}
