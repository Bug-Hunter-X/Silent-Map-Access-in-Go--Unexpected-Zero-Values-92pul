```go
package main

import (
    "fmt"
)

func main() {
    var m = make(map[int]int)
    m[1] = 10
    value, ok := m[2]
    if ok {
        fmt.Println("Value found:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```