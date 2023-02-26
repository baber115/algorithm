package day_one

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// https://leetcode.cn/problems/first-bad-version/

func firstBadVersionII(n int) int {
	return sort.Search(n, func(i int) bool {
		return isBadVersion(i)
	})
}

func firstBadVersionI(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return true
}
