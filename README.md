
# golang-algo-practice 

![golang-algo-practice-build](https://github.com/snghnaveen/golang-algo-practice/actions/workflows/push.yaml/badge.svg)

This repository includes a collection of algorithms and data structure programs implemented in Go. The problems are sourced from platforms such as `GeeksforGeeks`, `LeetCode`, `ChatGPT`, and others. Each program is written as Go test cases, ensuring they are fully functional and tested with multiple inputs.

### Running test cases

Prerequisite :
- go1.23
- Docker (optional)

Test (using Docker) :
- ``` docker-compose up```

---

Test (using go build) :
- ```make tests```

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
    <summary> Delete / pop last element in slice</summary>

```
s =  s[:len(s)-1]
```
</details>

---

<details>
    <summary> Delete / pop first element in slice</summary>

```
s =  s[1:]
```
</details>

---

<details>
    <summary> Add element in front slice</summary>

```
// add 122 at at index 0
m = slices.Insert(m, 0, 122)
```
or
```
temp = append(temp, 0)
temp = append([]int{0}, temp...)
```
</details>


---

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

---

<details>
    <summary> Split a string into a slice of substrings based on whitespace </summary>

```
str := "  Hello,   world!   This is   Go.  "

for i, v := range strings.Fields(str) {
    fmt.Println(i, v)
}

// 0 Hello,
// 1 world!
// 2 This
// 3 is
// 4 Go.
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
