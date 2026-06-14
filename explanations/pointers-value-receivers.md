### Intro
In Go you can use pointers to point to a specific memory address instead of making a copy of something etc. If you change something on pointer it will also be changed on the original (where it points to).

There are two common ways of using functions in Go - Pointer-Receivers versus Value-Receivers. They look similar but serve different purposes, here's an explanation to that.
### Code
**Data structure**
Let's start by defining a struct - its just a blueprint for a data structure and nothing special yet.
```Go
type model struct {
    id          int64
    name        string
    description string
}
```
**Function 1**
Now, we're creating two functions that compute an id based on this model. The first function is a **Pointer-Receiver** because it takes a pointer as input and uses it to compute the id.
```Go
func foo(m *model) {
    // business logic, which computes an id
    id := ...
    // change the value behind the pointer
    m.id = id
}
```
You can see that the function doesn't return something - this is optional because it directly modifies the original object we've defined first. You can see it on this code example here:
```Go
myModel := model{name: "foo"}
foo(&myModel)           // pass the ADDRESS with &
fmt.Println(myModel.id) // ✅ id has been set!
```
**Function 2**
The second function looks very similar to the first one if you're not familiar with Go but behaves very different, it's a **Value-Receiver** because it takes (in this case) just a value which is based on the data structure defined above.
```Go
func foo(m model) int64 {
    // business logic, which computes an id
    id := ...
    // return the id instead
    return id
}
```
If you use `model` without the `*` Go will copy the entire data structure and hands the function its private copy. Any changes to `m` inside the function are thrown away after its return, thats why we also need an explicit `return` here - otherwise we wouldn't be able to retrieve the computed id from it as the original structure stays untouched.
```Go
myModel := model{name: "foo"}
id := foo(myModel)      // pass the VALUE (a copy is made)
fmt.Println(myModel.id) // still 0 — unchanged
fmt.Println(id)         // ✅ you have the id here instead
```
### When to use which?
Both functions get the job done so this really depends on your individual use case. Pointer-Receivers mutate the original while Value-Receivers won't touch it. Therefore using Pointers, your function is more memory-efficient as it would be if you use Values. 

I could continue the list further but as a rule of thumb in Go:
- Use **pointers** if you want to **modify** the original state or if you data structure is large (to avoid copying large objects).
- Use **values** if the function is **read-only** and/or you want to make sure that nothing will touch your original data structure.