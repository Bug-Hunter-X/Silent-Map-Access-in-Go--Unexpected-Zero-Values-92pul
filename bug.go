```go
func main() {
    var m = make(map[int]int)
    m[1] = 10
    fmt.Println(m[2]) // This will print 0, not cause an error
}
```