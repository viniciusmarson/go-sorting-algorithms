// Package nlogn implements algorithm with complexity n log n
package nlogn

// MergeSort will order the list recived using the Merge Sort algorithm.
func MergeSort(list []int) {
	order(list, 0, len(list)-1)
}

func order(list []int, start int, end int) {
	if start < end && start >= 0 && end < len(list) && len(list) != 0 {
		meio := int((end + start) / 2)
		order(list, start, meio)
		order(list, meio+1, end)
		merge(list, start, meio, end)
	}
}

func merge(list []int, start int, meio int, end int) {

	var auxiliar = make([]int, len(list))

	for i := start; i <= end; i++ {
		auxiliar[i] = list[i]
	}

	i, j, k := start, meio+1, start

	for i <= meio && j <= end {
		if auxiliar[i] < auxiliar[j] {
			list[k] = auxiliar[i]
			i++
		} else {
			list[k] = auxiliar[j]
			j++
		}
		k++
	}

	for i <= meio {
		list[k] = auxiliar[i]
		i++
		k++
	}

	for j <= end {
		list[k] = auxiliar[j]
		j++
		k++
	}
}
