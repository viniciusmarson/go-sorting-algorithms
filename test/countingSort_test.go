package test

import (
	"fmt"
	"testing"

	"../n"
)

func TestCountingSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	response := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	n.CountingSort(response)

	if !AssertAreEqual(expected, response) {
		t.Error("The CountingSort algorithm is not working well")
		fmt.Println(response)
	}

}
