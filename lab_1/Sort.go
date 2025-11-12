package main

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, lowIndex int, highIndex int) {
	if lowIndex < highIndex {
		middleIndex := (lowIndex + highIndex) / 2
		mergeSort(arr, lowIndex, middleIndex)
		mergeSort(arr, middleIndex+1, highIndex)
		merge(arr, lowIndex, middleIndex, highIndex)
	}
}

func merge(arr []int, lowIndex int, middleIndex int, highIndex int) {
	left := lowIndex
	right := middleIndex + 1
	tempArray := make([]int, highIndex-lowIndex+1)
	index := 0

	for left <= middleIndex && right <= highIndex {
		if arr[left] < arr[right] {
			tempArray[index] = arr[left]
			left++
		} else {
			tempArray[index] = arr[right]
			right++
		}
		index++
	}

	for i := left; i <= middleIndex; i++ {
		tempArray[index] = arr[i]
		index++
	}
	for i := right; i <= highIndex; i++ {
		tempArray[index] = arr[i]
		index++
	}
	for i := range tempArray {
		arr[lowIndex+i] = tempArray[i]
	}
}
