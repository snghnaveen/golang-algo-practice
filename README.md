
# golang-algo-practice 

![golang-algo-practice-build](https://github.com/snghnaveen/golang-algo-practice/actions/workflows/push.yaml/badge.svg)

This repo contains various algorithms & data structure programs written in golang.
Questions are taken from various sources like `geekforgeeks`, `leetcode` , `ChatGPT`, etc.
Programs are written in form of `go test` cases. They are in fully working condition with multiple inputs.

### Running test cases

Prerequisite :
- go1.23
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
.
├── array
│   └── # questions realted to array
├── backtracking
├── basics
├── concurrency
├── internal  <----- For datastructure like linked list, stack etc.
├── json
├── linkedlist
├── matrix
├── queue
├── search
├── sorting
├── stack
├── string
└── tree
```


### Index
| Type / Category | Source                         | Content                            |
| --------------- | ------------------------------ | ---------------------------------- |
| `Array`         | [array](./array)               | [README](./array/README.md)        |
| `Backtracking`  | [backtracking](./backtracking) | [README](./backtracking/README.md) |
| `Basics`        | [basics](./basics)             | [README](./basics/README.md)       |
| `Concurrency`   | [concurrency](./concurrency)   | [README](./concurrency/README.md)  |
| `Json`          | [json](./json)                 | [README](./json/README.md)         |
| `LinkedList`    | [linkedlist](./linkedlist)     | [README](./linkedlist/README.md)   |
| `Matrix`        | [matrix](./matrix)             | [README](./matrix/README.md)       |
| `Stack`         | [stack](./stack)               | [README](./stack/README.md)        |
| `Queue`         | [queue](./queue/)              | [README](./queue/README.md)        |
| `Search`        | [search](./search)             | [README](./search/README.md)       |
| `Sorting`       | [sorting](./sorting)           | [README](./sorting/README.md)      |
| `String`        | [string](./string)             | [README](./string/README.md)       |
| `JSON Parsing`  | [json](./json)                 | [README](./json/README.md)         |
| `Tree`          | [tree](./tree/)                | [README](./tree/README.md)         |


### Cheatsheet

<details>
    <summary> Find middle element </summary>

```
mid := left + (right - left) / 2
```

</details>

---

<details>
    <summary> String to rune </summary>

```
runes := []rune(str)
```

</details>

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
- https://leetcode.com/problems/merge-intervals/description/
- https://leetcode.com/problems/zigzag-conversion/
- https://leetcode.com/problems/string-compression/description/
- https://leetcode.com/problems/furthest-building-you-can-reach/description/
- https://leetcode.com/problems/integer-to-roman/description/?envType=study-plan-v2&envId=top-interview-150
- https://leetcode.com/problems/even-odd-tree/description/
- https://www.geeksforgeeks.org/find-a-triplet-that-sum-to-a-given-value/
- https://leetcode.com/problems/merge-nodes-in-between-zeros/submissions/1309505523/
