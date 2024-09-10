# This is how to declare a method following documentation of go:

[The documentation about Methods](https://go.dev/wiki/MethodSets)

> A [code snippet](https://go.dev/play/p/DUnNYS-v81y) that I've made for tests purpose.
> You can add return to the method like in another langs too.

```go
package main

import "fmt"

type User struct {
    name string
    age int
}

func NewUser(NewName string, NewAge int) *User {
    return &User {
        name: NewName,
        age: NewAge,
    }
}

func (u \*User) PrintInfo() {
    fmt.Println("Name:" + u.name)
    fmt.Println("Yo:", u.age)
}

func main() {
    user := NewUser("Stete", 24)
    user.PrintInfo()
}
```
