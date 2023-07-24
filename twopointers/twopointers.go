package twopointers

func LengthOfUniques(numbers []int) int {
	slow := 0
	length := len(numbers)

	for fast := 0; fast < length; fast++ {
		if numbers[fast] != numbers[slow] {
			slow++
			numbers[slow] = numbers[fast]
		}
	}

	return slow + 1
}

func TwoSumSorted(numbers []int, target int) [2]int {
	startingPtr := 0
	endingPtr := len(numbers) - 1
	sum := numbers[startingPtr] + numbers[endingPtr]

	for sum != target {
		if sum > target {
			endingPtr--
		} else if sum < target {
			startingPtr++
		}
		sum = numbers[startingPtr] + numbers[endingPtr]
	}

	return [2]int{startingPtr, endingPtr}
}
