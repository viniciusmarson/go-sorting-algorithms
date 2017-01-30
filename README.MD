# Sorting algorithms

This repository is only for educational purposes. You will find the most used sorting algorithms here implemented in the [Go Programming Language](https://golang.org/)

##### Complexity n²
- Bubble Sort [DONE]
- Insection Sort [DONE]
- Selection Sort [DONE] 

##### Complexity n log n 
- Merge Sort [DONE]
- Quicksort [DONE]
- Heapsort [TODO]

##### Complexity n
- Counting Sort [DONE]
- Bucket Sort [TODO]
- Radix [TODO]

### Running

Requires [Go](https://golang.org/doc/install) to run.

In a terminal window with the GOPATH environment variable set to the GOPATH you want to work on, launch code.

If you are using Visual Studio Code you can follow this [tutorial](https://github.com/Microsoft/vscode-go) to configure him.

To execute the algorithms you just need to run the index file in the prompt. 

Example:

```sh
$ go run main.go
```

### Tests

Install [go-expect](https://github.com/viniciusmarson/go-expect):

```sh
$ go get github.com/viniciusmarson/go-expect/expect
```

In a terminal window run the command:

```sh
$ go test -v
```

Example:

```sh
C:\dev\github\go-sorting-algorithms\n2 (master)
λ go test -v
=== RUN   TestBubbleSort
--- PASS: TestBubbleSort (0.00s)
=== RUN   TestInsertionSort
--- PASS: TestInsertionSort (0.00s)
=== RUN   TestSelectionSort
--- PASS: TestSelectionSort (0.00s)
PASS
ok      _/C_/dev/github/go-sorting-algorithms/n2        0.398s
```

![result.png](http://image.prntscr.com/image/505cf7287a434380b5e21d3aedb2fdda.png)

![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)