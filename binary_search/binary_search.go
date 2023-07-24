package binarysearch

func binarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	mid := end / 2

	for start <= end {
		switch {
		case target == arr[mid]:
			return mid
		case target > arr[mid]:
			start = mid + 1
		case target < arr[mid]:
			end = mid - 1
		}
		mid = (start + end) / 2
	}

	return -1
}

// The findBoundary function returns the index of the first true
// element in the sorted Boolean slice `arr`. If `arr` does not contain a
// true element, the function returns `-1`.
func findBoundary(arr []bool) int {
	firstP := 0
	lastP := len(arr) - 1
	midP := lastP >> 1

	for firstP <= lastP {
		switch {
		case arr[midP] && midP == 0:
			return midP
		case arr[midP] && !arr[midP-1]:
			return midP
		case arr[midP]:
			lastP = midP
		case !arr[midP]:
			firstP = midP + 1
		}
		midP = (firstP + lastP) >> 1
	}

	return -1
}

// The findBoundaryShorter function is an alternative implementation of
// find Boundary and returns the index of the first true element in the
// sorted Boolean slice `arr`. If `arr` does not contain a true element,
// the function returns `-1`.
func findBoundaryShorter(arr []bool) int {
	firstP := 0
	lastP := len(arr) - 1
	boundary := -1

	for lastP >= firstP {
		midP := (firstP + lastP) >> 1
		if arr[midP] {
			boundary = midP
			lastP = boundary - 1
		} else {
			firstP = midP + 1
		}
	}

	return boundary
}

func firstNotSmaller(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	boundary := -1

	for left <= right {
		mid := (left + right) >> 1
		if arr[mid] >= target {
			boundary = mid
			right = boundary - 1
		} else {
			left = mid + 1
		}
	}

	return boundary
}

func findFirstOccurrence(arr []int, target int) int {
	boundary := -1
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) >> 1
		switch {
		case arr[mid] == target:
			boundary = mid
			right = boundary - 1
		case arr[mid] < target:
			left = mid + 1
		case arr[mid] > target:
			right = mid - 1
		}
	}

	return boundary
}

// The squareRoot function returns the closest floored integer <= to the
// square root of a given integer `n`. A return value of `-1` means
// that the function failed to find a floored square root of `n`.
func squareRoot(n int) int {
	left := 0
	right := n
	closest := -1

	for left <= right {
		mid := (left + right) >> 1
		product := mid * mid

		if n < product {
			// Our guess is too high, discard everything above mid
			right = mid - 1
		} else {
			// Our guess is either exactly correct or our best match so far,
			// discard everything else below mid
			closest = mid
			left = mid + 1
		}
	}

	return closest
}

// The findMinRotated function returns the smallest int of a slice of
// ints. The slice `arr` should be an ordered slice that has been rotated
// so that the smallest element is no longer the first element.
func findMinRotated(arr []int) int {
	left := 0
	right := len(arr) - 1
	end := arr[right]
	boundary := -1

	for left <= right {
		mid := (left + right) >> 1

		if arr[mid] <= end {
			right = mid - 1
			boundary = mid
		} else {
			left = mid + 1
		}
	}

	return boundary
}

// The peakOfMoutainArray function returns the index of the highest value
// in slice `arr`. A mountain array should have at least 3 elements. The
// "peak" of the mountain is the highest value at index `k`. Elements
// preceding `k` should increase up to `k`. Conversely, elements following
// `k` should decrease in size.
func peakOfMountainArray(arr []int) int {
	rightEdge := len(arr) - 1
	left := 0
	right := rightEdge
	boundary := -1

	for left <= right {
		mid := (left + right) >> 1
		if mid == rightEdge || arr[mid] > arr[mid+1] {
			// at the end of the slice - peak is here or to the left
			boundary = mid
			right = mid - 1
		} else {
			// the peak must be to the right if we get here
			left = mid + 1
		}
	}

	return boundary
}

// The newspapersSplit function returns the amount if time it would take
// `numCoworkers` to consume `newspapersReadTimes`. In abstract, this
// function distributes the newspapers to coworkers as evenly as possible
// and returns the sum of the largest subslice by value.
func newspapersSplit(newspapersReadTimes []int, numCoworkers int) int {
	var minTime, maxTime int
	ans := -1

	// The largest single read time is the lowest our answer can be
	for _, i := range newspapersReadTimes {
		if i > minTime {
			minTime = i
		}
		maxTime++
	}

	for minTime <= maxTime {

	}

	return ans
}

func newspapersSplitFeasible(newspapersReadTimes []int, numCoworkers int, limit int) {
	var time, numWorkers int

	for _, readTime := range newspapersReadTimes {
		if time+readTime > limit {
			time = 0
			numWorkers++
		}

		time += readTime
	}

	// Edge case, we still have more time to account for and need another coworker
	if time != 0 {
		numWorkers++
	}
}
