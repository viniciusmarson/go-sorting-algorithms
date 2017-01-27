package test

func AssertAreEqual(expected []int, result []int) bool {
	equals := true

	for i := range expected {
		if expected[i] != result[i] {
			equals = false
		}
	}

	return equals
}
