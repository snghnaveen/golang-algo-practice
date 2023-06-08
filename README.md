# golang-algo-practice 

![golang-algo-practice-build](https://github.com/snghnaveen/golang-algo-practice/actions/workflows/push.yaml/badge.svg)

This repo contains various algorithms & data structure programs written in golang.
Questions are taken from various sources like `geekforgeeks`, `leetcode` , `ChatGPT`, etc.
Programs are written in form of `go test` cases. They are in fully working condition with multiple inputs.

### Running test cases

Prerequisite :
- go1.20
- Docker (optional)

Test (Docker) :
- ``` docker-compose up```

Test (build) :

- ```make tests```
OR
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
| `LinkedList`   | [linkedlist](./linkedlist) | [README](./linkedlist/README.md)
| `Stack`   | [stack](./stack) | [README](./stack/README.md)
| `Search`   | [search](./search) | [README](./search/README.md)
| `Sorting`   | [sorting](./sorting) | [README](./sorting/README.md)
| `JSON Parsing`   | [json](./json) | [README](./json/README.md)


### TODO
- https://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/
- https://www.geeksforgeeks.org/efficiently-implement-k-stacks-single-array/
- Find the longest palindrome in a string
- Given an array of integers, print all combinations of size X. (merged)
- Given an array of integers A, print all its subsets. (merged)
- Maze questions
- https://www.geeksforgeeks.org/sort-an-array-of-0s-1s-and-2s/
- Longest Increasing Subsequence (LIS)
- https://github.com/ruppysuppy/Daily-Coding-Problem-Solutions
- https://leetcode.com/problems/number-of-islands/