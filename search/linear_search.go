/*
* Linear search -
* Walks the entire array until the needle is found.
* Worst case time complexity: O(n) for an array of length n.
*/

package search

func LinearSearch(haystack []int, needle int) int {
	for i, item := range haystack {
		if item == needle {
			return i
		}
	}

	return -1
}