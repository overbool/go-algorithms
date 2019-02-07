## interface

```go
type Stack interface {
    IsEmpty() bool
    Len() int
    Push(interface{})
    Pop() interface{}
    Top() interface{}
}
```