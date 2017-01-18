// Package n2 implements algorithm with complexity n²
package n2

// SelectionSort will order the list recived using the Selection Sort algorithm.
func SelectionSort(list []int) {
	for i := 0; i < cap(list); i++ {
		min := i

		for j := i + 1; j < len(list); j++ {
			if list[j] < list[min] {
				min = j
			}
		}

		if min != i {
			aux := list[min]
			list[min] = list[i]
			list[i] = aux
		}
	}
}
