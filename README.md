# golang-algo-practice
This repo contains algorithm programs in golang.

Algorithm are written in form of test cases.

## Run
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

## Directory structure
```
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── internal <----- For datastructure like linked list, stack etc.
├── main.go
└── string
    └── 
└── array
    └── 

```


### Index
| Question | Type/ Tag  | Source |
| --------| ------------- | ------------- |
| Check if given strings are rotations of each other or not | `String`  | [string_rotation_test.go](./string/string_rotation_test.go)  |
