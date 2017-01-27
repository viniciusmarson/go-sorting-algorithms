package main

import (
	"fmt"

	"./n2"

	"./nlogn"

	"./n"
)

func main() {
	fmt.Println("Complexity n²")
	bubble()
	insertion()
	selection()

	fmt.Println("Complexity n log n")
	merge()
	quick()

	fmt.Println("Complexity n")
	counting()
}

func bubble() {
	list := []int{20, 10, 13, 14, 15, 80, 1, 69, 95, 44, 33}
	fmt.Println("Bubble Sort algorithm:")
	fmt.Print("Infromed list: ")
	fmt.Println(list)

	n2.BubbleSort(list)

	fmt.Print("Ordened list: ")
	fmt.Println(list)

	fmt.Println("")
}

func insertion() {
	list := []int{20, 10, 13, 14, 15, 80, 1, 69, 95, 44, 33}
	fmt.Println("Insertion Sort algorithm:")
	fmt.Print("Infromed list: ")
	fmt.Println(list)

	n2.InsertionSort(list)

	fmt.Print("Ordened list: ")
	fmt.Println(list)

	fmt.Println("")
}

func selection() {
	list := []int{20, 10, 13, 14, 15, 80, 1, 69, 95, 44, 33}
	fmt.Println("Selection Sort algorithm:")
	fmt.Print("Infromed list: ")
	fmt.Println(list)

	n2.SelectionSort(list)

	fmt.Print("Ordened list: ")
	fmt.Println(list)

	fmt.Println("")
}

func merge() {
	list := []int{20, 10, 13, 14, 15, 80, 1, 69, 95, 44, 33}
	fmt.Println("Merge Sort algorithm:")
	fmt.Print("Infromed list: ")
	fmt.Println(list)

	nlogn.MergeSort(list)

	fmt.Print("Ordened list: ")
	fmt.Println(list)

	fmt.Println("")
}

func quick() {
	list := []int{20, 10, 13, 14, 15, 80, 1, 69, 95, 44, 33}
	fmt.Println("Quick Sort algorithm:")
	fmt.Print("Infromed list: ")
	fmt.Println(list)

	nlogn.QuickSort(list)

	fmt.Print("Ordened list: ")
	fmt.Println(list)

	fmt.Println("")
}

func counting() {
	list := []int{20, 10, 13, 14, 15, 80, 1, 69, 95, 44, 33}
	fmt.Println("Counting Sort algorithm:")
	fmt.Print("Infromed list: ")
	fmt.Println(list)

	n.CountingSort(list)

	fmt.Print("Ordened list: ")
	fmt.Println(list)

	fmt.Println("")
}
