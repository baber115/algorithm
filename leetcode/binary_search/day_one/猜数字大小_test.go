package day_one

import "testing"

// https://leetcode.cn/problems/guess-number-higher-or-lower/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=f2ex1x6

func TestGuessNumber(t *testing.T) {

}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	slow, fast := 1, n
	for slow <= fast {
		mid := slow + (fast-slow)/2
		guessAns := guess(mid)
		if guessAns == -1 {
			fast = mid - 1
		} else if guessAns == 1 {
			slow = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func guess(num int) int {
	return 1
}
