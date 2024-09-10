> Exclusivelly in this project will have some comments inside of functions
> but it's not my practice comment inside of a function or struct

# Comments of struct should follow this pattern:

- What the struct represents
- What components the struct have

> And realize that have a `type` keyword. You can read more about `type definition` and `type identity` in [golang documentation](https://go.dev/doc/)

```go
// UserInfo represents data of a user.
//
// It contains essential data to build a user instance.
type UserInfo struct {
    name string // here can have some explanation about field
    age int
    country string
}
```

# Comments for functions:

- What the function does
- Parameters section
- Return sectionn

```go
// NewUser creates and returns a new instance of UserInfo
//
// It initializes an UserInfo struct with the provided parameters
//
// Parameters:
//
// name string: Name of the user.
// age int: Age of user.
// country string: Country of the user.
//
// Returns:
//
//  *UserInfo: A pointer to a newly created UserInfo instance.
func NewUser(name string, age int, country string) *UserInfo {
    return &UserInfo{
        name: name,
        age: age,
        country: country,
    }
}
```

---

For a while that's all folks... If have some updates I'll back here
