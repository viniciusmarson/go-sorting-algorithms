package test

import (
	"fmt"
	"testing"

	"../n2"
)

func TestBubbleSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	response := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	n2.BubbleSort(response)

	if !AssertAreEqual(expected, response) {
		t.Error("The BubbleSort algorithm is not working well")
		fmt.Println(response)
	}

}
