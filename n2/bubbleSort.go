// Package n2 implements algorithm with complexity n²
package n2

// BubbleSort will order the list recived using the Bubble Sort algorithm.
func BubbleSort(list []int) {
	changed := true

	for changed {

		changed = false

		for position := 0; position <= len(list)-2; position++ {
			if list[position] > list[position+1] {
				var temp = list[position]
				list[position] = list[position+1]
				list[position+1] = temp
				changed = true
			}
		}
	}
}
