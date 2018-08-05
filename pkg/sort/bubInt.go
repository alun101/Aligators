package bubint

//Asc recieves an int slice and bubble sorts it in to ascending order
func Asc(numbers []int) {
	length := len(numbers)
	previousPasses := 0

	for outer := 0; outer < (length - 1); outer++ {
		firstIndex := 0
		secondIndex := 1
		swap := false
		for inner := 0; inner < ((length - 1) - previousPasses); inner++ {
			if numbers[firstIndex] > numbers[secondIndex] {
				numbers[firstIndex], numbers[secondIndex] = numbers[secondIndex], numbers[firstIndex]
				swap = true
			}
			firstIndex++
			secondIndex++
		}
		if !swap {
			return
		}
		previousPasses++
	}
}

//Des recieves an int slice and bubble sorts it in to descending order
func Des(numbers []int) {
	length := len(numbers)
	previousPasses := 0

	for outer := 0; outer < (length - 1); outer++ {
		firstIndex := 0
		secondIndex := 1
		swap := false
		for inner := 0; inner < ((length - 1) - previousPasses); inner++ {
			if numbers[firstIndex] < numbers[secondIndex] {
				numbers[firstIndex], numbers[secondIndex] = numbers[secondIndex], numbers[firstIndex]
				swap = true
			}
			firstIndex++
			secondIndex++
		}
		if !swap {
			return
		}
		previousPasses++
	}
}
