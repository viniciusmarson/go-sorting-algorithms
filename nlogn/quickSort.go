// Package nlogn implements algorithm with complexity n log n
package nlogn

// QuickSort will order the list recived using the Quick Sort algorithm.
func QuickSort(list []int) {
	quick(list, 0, len(list)-1)
}

func quick(elements []int, start int, end int) {
	if start < end {
		var pivot = partition(elements, start, end)
		quick(elements, start, pivot-1)
		quick(elements, pivot+1, end)
	}
}

func partition(elements []int, start int, end int) int {

	pivot, i := elements[end], start

	for j := start; j <= end-1; j++ {
		if elements[j] <= pivot {
			aux := elements[i]
			elements[i] = elements[j]
			elements[j] = aux
			i++
		}
	}

	aux := elements[i]
	elements[i] = elements[end]
	elements[end] = aux

	return i
}
