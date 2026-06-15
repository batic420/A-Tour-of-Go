# A Tour of Go

Personal notes and exercises from working through [A Tour of Go](https://go.dev/tour/), the official interactive introduction to the Go programming language.

## What's here

- **Concept notes** — explanations of core Go concepts encountered throughout the tour
- **Exercises** — solutions to the coding exercises in `exercises/`

## Key Concepts

### Pointers

A pointer holds the memory address of a value. Go uses `*` to declare a pointer type and dereference it, and `&` to take the address of a value.

```go
i := 42
p := &i       // p is a *int pointing to i
fmt.Println(*p) // read i through the pointer → 42
*p = 21       // set i through the pointer
fmt.Println(i)  // → 21
```

Unlike C, Go has no pointer arithmetic.

### Slices

A slice is a dynamically-sized view into an array. The built-in `make` function creates a slice with a given length (and optional capacity).

```go
s := make([]int, 5)       // len=5, cap=5, all zeros
t := make([][]uint8, dy)  // 2-D slice
```

Slices do not copy data — multiple slices can reference the same underlying array.

### Structs

A struct groups fields together. Fields are accessed with `.` and can be reached through a pointer directly (no need to dereference manually).

```go
type Vertex struct {
    X, Y int
}

v := Vertex{1, 2}
p := &v
p.X = 10  // shorthand for (*p).X
```

### Interfaces

An interface is a set of method signatures. A type implements an interface implicitly — no `implements` keyword needed.

```go
type Stringer interface {
    String() string
}
```

Any type with a `String() string` method satisfies `Stringer`.

### Goroutines & Channels

Goroutines are lightweight threads managed by the Go runtime. Channels provide a typed pipe for goroutines to communicate safely.

```go
go func() { /* runs concurrently */ }()

ch := make(chan int)
go func() { ch <- 42 }()
v := <-ch  // receive
```

## Exercises

| File | Tour exercise |
|------|--------------|
| [exercises/slices/slices.go](exercises/slices/slices.go) | [Slices](https://go.dev/tour/moretypes/18) — generate a 2-D image using a custom pixel calculation |
| [exercises/maps/maps.go](exercises/maps/maps.go) | [Maps](https://go.dev/tour/moretypes/23) — create a word counter for a string using map generation |
| [exercises/fibonacci/fibonacci.go](exercises/fibonacci/fibonacci.go) | [Maps](https://go.dev/tour/moretypes/26) — generate a list of fibonacci numbers using a function closure |

## Running the exercises

```bash
cd exercises
go run ./subdirectoy # Replace with directory name for the exercise you want to run
```

Some exercises (like `slices.go`) depend on `golang.org/x/tour`, which is already declared in `go.mod`.
