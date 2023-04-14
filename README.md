# golang-algo-practice
This repo contains algorithm programs in golang.

Algorithm are written in form of test cases.

### Running test cases

Prerequisite :
- go1.20 

Test :
- Run all tests 
    ```
    go test -v ./...
    ```
- Run specific test
    ```
    go test -v -run TestCheckStringAreRotations ./...
    ```

- Clear test cache 
    ```
    go clean -testcache
    ```

### Directory structure
```
├── Makefile
├── README.md
├── array
│   └── 
├── go.mod
├── go.sum
├── internal <----- For datastructure like linked list, stack etc.
├── main.go
└── string
    └── 
```


### Index
| Type / Category  |     Source         | Content   |
| ---------- |       -------      | -----  |
| `String`   | [string](./string) | [README](./string/README.md)
| `Array`   | [array](./array) | [README](./array/README.md)
| `LinkedList`   | [array](./linkedlist) | [README](./linkedlist/README.md)


### TODO
- https://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/
- https://www.geeksforgeeks.org/efficiently-implement-k-stacks-single-array/
- Find the longest palindrome in a string
- Given an array of integers, print all combinations of size X. (merged)
- Given an array of integers A, print all its subsets. (merged)
- Maze questions
- https://www.geeksforgeeks.org/sort-an-array-of-0s-1s-and-2s/
