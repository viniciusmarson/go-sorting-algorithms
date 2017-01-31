package n2

import (
	"testing"

	"github.com/viniciusmarson/go-expect/expect"
)

func TestSelectionSort(t *testing.T) {
	theExpectedResponse := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	response := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(response)
	expect.This(response).ToBe(theExpectedResponse)(t)
}
