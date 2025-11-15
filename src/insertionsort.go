package main

// insertionSort performs an in place insertion sort on the array
func insertionSort(arr *[]pair) {
	for i := 1; i < len(*arr); i++ {
		pair := (*arr)[i]
		j := i - 1

		for {
			if !(j >= 0 && pair.value > (*arr)[j].value) {
				break
			}

			(*arr)[j+1] = (*arr)[j]
			j--
		}

		(*arr)[j+1] = pair
	}
}
