// Package n2 implements algorithm with complexity n²
package n2

// InsertionSort will order the list recived using the Insertion Sort algorithm.
func InsertionSort(list []int) {
	for i, content := range list {

		key := content
		j := i

		for (j > 0) && (list[j-1] > key) {
			list[j] = list[j-1]
			j--
		}

		list[j] = key
	}
}
