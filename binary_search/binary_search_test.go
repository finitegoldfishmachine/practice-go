package binarysearch

import "testing"

func TestVanilla(t *testing.T) {
	testCases := []struct {
		desc   string
		value  []int
		target int
		want   int
	}{
		{
			desc:   "sequential",
			value:  []int{1, 2, 3, 4, 5, 6, 7},
			target: 5,
			want:   4,
		},
		{
			desc:   "gaps",
			value:  []int{1234, 2345, 23456, 2346314, 1123412341234, 12348679122348},
			target: 23456,
			want:   2,
		},
		{
			desc:   "small amount even",
			value:  []int{1234, 2345, 23456, 123123489123},
			target: 123123489123,
			want:   3,
		},
		{
			desc:   "small amount odd",
			value:  []int{1234, 2345, 23456},
			target: 1234,
			want:   0,
		},
		{
			desc:   "target is last",
			value:  []int{1234, 2345, 23456, 2346314, 1123412341234, 12348679122348},
			target: 12348679122348,
			want:   5,
		},
		{
			desc:   "single element",
			value:  []int{1234},
			target: 1234,
			want:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := binarySearch(tC.value, tC.target)
			if tC.want != got {
				t.Fatalf("%v - got %v want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestDoesntExist(t *testing.T) {
	arr := []int{
		1234,
		2345,
		23456,
		2346314,
		1123412341234,
		12348679122348,
		123486791223481234,
		1234867912234812341,
	}

	testCases := []struct {
		desc   string
		value  []int
		target int
		want   int
	}{
		{
			desc:   "target doesn't exist mid",
			value:  arr,
			target: 123482134,
			want:   -1,
		},
		{
			desc:   "target doesn't exist first",
			value:  arr,
			target: 123,
			want:   -1,
		},
		{
			desc:   "target doesn't exist last",
			value:  arr,
			target: 123486791223485,
			want:   -1,
		},
		{
			desc:   "no elements",
			value:  []int{},
			target: 1234,
			want:   -1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := binarySearch(tC.value, tC.target)
			if tC.want != got {
				t.Fatalf("%v - got %v want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestFindBoundary(t *testing.T) {
	testCases := []struct {
		desc  string
		value []bool
		want  int
	}{
		{
			desc:  "Simple case",
			value: []bool{false, false, true, true, true},
			want:  2,
		},
		{
			desc:  "Completely left",
			value: []bool{true, true, true, true, true},
			want:  0,
		},
		{
			desc:  "Completely right",
			value: []bool{false, false, false, false, true},
			want:  4,
		},
		{
			desc:  "Doesn't exist",
			value: []bool{false, false, false, false, false},
			want:  -1,
		},
		{
			desc:  "Single element true",
			value: []bool{true},
			want:  0,
		},
		{
			desc:  "Single element false",
			value: []bool{false},
			want:  -1,
		},
		{
			desc:  "No elements",
			value: []bool{},
			want:  -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := findBoundary(tC.value)
			if tC.want != got {
				t.Fatalf("%v - got %v want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestFindBoundaryShorter(t *testing.T) {
	testCases := []struct {
		desc  string
		value []bool
		want  int
	}{
		{
			desc:  "Simple case",
			value: []bool{false, false, true, true, true},
			want:  2,
		},
		{
			desc:  "Completely left",
			value: []bool{true, true, true, true, true},
			want:  0,
		},
		{
			desc:  "Completely right",
			value: []bool{false, false, false, false, true},
			want:  4,
		},
		{
			desc:  "Doesn't exist",
			value: []bool{false, false, false, false, false},
			want:  -1,
		},
		{
			desc:  "Single element true",
			value: []bool{true},
			want:  0,
		},
		{
			desc:  "Single element false",
			value: []bool{false},
			want:  -1,
		},
		{
			desc:  "No elements",
			value: []bool{},
			want:  -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := findBoundaryShorter(tC.value)
			if tC.want != got {
				t.Fatalf("%v - got %v want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestFirstNotSmaller(t *testing.T) {
	testCases := []struct {
		desc   string
		arr    []int
		target int
		want   int
	}{
		{
			desc:   "Simple case",
			arr:    []int{1, 2, 3, 4, 5},
			target: 2,
			want:   1,
		},
		{
			desc:   "Completely left",
			arr:    []int{2, 3, 4, 5, 6},
			target: 1,
			want:   0,
		},
		{
			desc:   "Completely right",
			arr:    []int{2, 2, 3, 3, 4},
			target: 4,
			want:   4,
		},
		{
			desc:   "Doesn't exist",
			arr:    []int{2, 2, 2},
			target: 3,
			want:   -1,
		},
		{
			desc:   "Single element true",
			arr:    []int{2},
			target: 1,
			want:   0,
		},
		{
			desc:   "Single element false",
			arr:    []int{2},
			target: 3,
			want:   -1,
		},
		{
			desc:   "No elements",
			arr:    []int{},
			target: 0,
			want:   -1,
		},
		{
			desc:   "Same values",
			arr:    []int{0, 0},
			target: 0,
			want:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := firstNotSmaller(tC.arr, tC.target)
			if tC.want != got {
				t.Fatalf("%v - got %v want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestFindFirstOccurence(t *testing.T) {
	testCases := []struct {
		desc   string
		arr    []int
		target int
		want   int
	}{
		{
			desc:   "Simple case",
			arr:    []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			desc:   "Completely left",
			arr:    []int{2, 3, 4, 5, 6},
			target: 2,
			want:   0,
		},
		{
			desc:   "Completely right",
			arr:    []int{2, 2, 3, 3, 4},
			target: 4,
			want:   4,
		},
		{
			desc:   "Doesn't exist",
			arr:    []int{2, 2, 2},
			target: 3,
			want:   -1,
		},
		{
			desc:   "Single element true",
			arr:    []int{2},
			target: 2,
			want:   0,
		},
		{
			desc:   "Single element false",
			arr:    []int{2},
			target: 3,
			want:   -1,
		},
		{
			desc:   "No elements",
			arr:    []int{},
			target: 0,
			want:   -1,
		},
		{
			desc:   "Several matches",
			arr:    []int{0, 1, 1, 1, 1},
			target: 1,
			want:   1,
		},
		{
			desc:   "Same values",
			arr:    []int{0, 0},
			target: 0,
			want:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := findFirstOccurrence(tC.arr, tC.target)
			if tC.want != got {
				t.Fatalf("%v - got %v want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	testCases := []struct {
		desc   string
		number int
		want   int
	}{
		{
			desc:   "Simple case",
			number: 9,
			want:   3,
		},
		{
			desc:   "Floored case",
			number: 10,
			want:   3,
		},
		{
			desc:   "28 floored float",
			number: 28,
			want:   5,
		},
		{
			desc:   "0",
			number: 0,
			want:   0,
		},
		{
			desc:   "1",
			number: 1,
			want:   1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := squareRoot(tC.number)
			if tC.want != got {
				t.Fatalf("%v - got %v want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestFindMinRotated(t *testing.T) {
	testCases := []struct {
		desc  string
		input []int
		want  int // `want` is the index of the smallest element
	}{
		{
			desc:  "basic case",
			input: []int{30, 40, 50, 10, 20},
			want:  3,
		},
		{
			desc:  "larger basic case",
			input: []int{3, 5, 7, 11, 13, 17, 19, 2},
			want:  7,
		},
		{
			desc:  "smallest in first half",
			input: []int{19, 2, 3, 5, 7, 11, 13, 17},
			want:  1,
		},
		{
			desc:  "exactly middle",
			input: []int{11, 13, 3, 5, 7},
			want:  2,
		},
		{
			desc:  "single element",
			input: []int{3},
			want:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := findMinRotated(tC.input)
			if tC.want != got {
				t.Fatalf("%v findMinRotated(%#v)- got %v want %v", tC.desc, tC.input, got, tC.want)
			}
		})
	}
}

func TestPeakOfMountainArray(t *testing.T) {
	testCases := []struct {
		desc  string
		input []int
		want  int
	}{
		{
			desc:  "basic case",
			input: []int{0, 1, 2, 3, 2, 1, 0},
			want:  3,
		},
		{
			desc:  "peak second element",
			input: []int{0, 10, 3, 2, 1, 0},
			want:  1,
		},
		{
			desc:  "completely left",
			input: []int{10, 3, 2, 1, 0},
			want:  0,
		},
		{
			desc:  "completely right",
			input: []int{0, 1, 2, 10},
			want:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := peakOfMountainArray(tC.input)
			if tC.want != got {
				t.Fatalf("%v peakOfMountainArray(%#v)- got %v want %v", tC.desc, tC.input, got, tC.want)
			}
		})
	}
}

func TestNewspapersSplit(t *testing.T) {
	testCases := []struct {
		desc         string
		readTimes    []int
		numCoworkers int
		want         int
	}{
		{
			desc:         "two readers",
			readTimes:    []int{7, 2, 5, 10, 8},
			numCoworkers: 2,
			want:         18,
		},
		{
			desc:         "three readers",
			readTimes:    []int{2, 3, 5, 7},
			numCoworkers: 3,
			want:         7,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := newspapersSplit(tC.readTimes, tC.numCoworkers)
			if tC.want != got {
				t.Fatalf("%v newspapersSplit(%#v, %#v)- got %v want %v", tC.desc, tC.readTimes, tC.numCoworkers, got, tC.want)
			}
		})
	}
}
