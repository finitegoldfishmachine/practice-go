package sort

type SortByInt []int

func (a SortByInt) Len() int           { return len(a) }
func (a SortByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByInt) Less(i, j int) bool { return a[i] < a[j] }

func SortList(nums SortByInt) []int {
	i := 1

	for i < nums.Len() {
		if nums.Less(i, i-1) {
			nums.Swap(i, i-1)
			if i != 1 {
				i--
			} else {
				i++
			}
		} else {
			i++
		}
	}

	return nums
}

func InsertionSort(nums SortByInt) []int {
	// pA tracks the element being sorted
	// pB tracks the highest index pA was at
	pA, pB := 1, 1
	sorting := false

	for pA < nums.Len() {
		if nums.Less(pA, pA-1) {
			if !sorting {
				sorting = true
			}
			nums.Swap(pA, pA-1)
			// The first element cannot be sorted, continue otherwise
			if pA > 1 {
				pA--
				continue
			}
		}
		if sorting {
			sorting = false
		}
		// Element is sorted, continue to next element
		pB++
		pA = pB
	}
	return nums
}

func InsertionSortCondensed(a SortByInt) []int {
	for i := 1; i < a.Len(); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a.Swap(j, j-1)
		}
	}
	return a
}
