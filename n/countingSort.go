// Package n implements algorithm with complexity n
package n

// CountingSort will order the list recived using the Couting Sort algorithm.
func CountingSort(v []int) {

	higher := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > higher {
			higher = v[i]
		}
	}

	c := make([]int, higher)

	for i := 0; i < len(v); i++ {
		c[v[i]-1]++
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	b := make([]int, len(v))
	for i := 0; i < len(b); i++ {
		b[c[v[i]-1]-1] = v[i]
		c[v[i]-1]--
	}

	for i := 0; i < len(b); i++ {
		v[i] = b[i]
	}

}
