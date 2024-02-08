/*
* Binary search -
* Given a sorted array, jump to the middle of the array and determine its relationship with the needle.
* Worse case time complexity: O(lg n), where n is the length of the array and lg is the base 2 notation of the logarithm.
*
* Explanation -
* arr [ ... | ... ]
*     0    N/2    N
* 
* Having the array reduces the search space by half. We keep halving until we reach a sub array of unit length in the worst case.
* N/2 --> N/4 --> N/8 --> ... --> N/(2^k) such that N/(2^k) = 1
* Therefore, k = lg(N)
*
* So if we do k jumps in the worst case,
* complexity = O(lg N)
*/

package search

func BinarySearch(haystack []int, needle int) int {
	low := 0
	high := len(haystack)

	for ; low < high; {
		mid := low + ((high - low) / 2)
		item := haystack[mid]

		if item == needle {
			return mid
		} else if item < needle {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return -1
}
