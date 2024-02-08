/*
* Bubble sort -
* An array a is considered to be sorted if for any ith element in the array, 
* where 0 <= i < N, a[i] <= a[i + 1]
* 
* Worst case time complexity: O(N ^ 2)
*/

package sort

func BubbleSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}