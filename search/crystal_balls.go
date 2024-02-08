/*
* Crystal ball problem -
* Given *two* crystal balls that will break if dropped from high enough
* distance, determine the exact spot in which it will break in the
* most optimized way.
*
* Framing in terms of a search problem -
* The verdict on whether the ball will break can be viewed as an array of verdicts.
* [ false, false, false, ... true, true, ... ]
* 0                          i               N
* 
* Where N is the highest possible height and i is the least height at which the balls break.
* Now, we just need to find the first occurance of "true" in an optimized way.
*
* The array is sorted. We can jump by some distance, call it d until the first ball breaks.
* Then we just walk the distance d from the previous point to the current point of breakage.
*
* Assuming d is sqrt(N) for an array of N
* worst case time complexity:
* N / sqrt(N) + sqrt(N) = sqrt(N) + sqrt(N) = 2 * sqrt(N)
* O(sqrt(N))
*/

package search

import "math"

func CrystalBalls(breaks []bool) int {
	distance := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	curr := 0
	lookAhead := 0
	for ;; {
		curr = lookAhead
		lookAhead += distance
		if lookAhead >= len(breaks) || breaks[lookAhead] {
			break
		}
	}

	
	i := curr
	for ; i <= lookAhead; i++ {
		if breaks[i] {
			return i
		}
	}

	return -1
}