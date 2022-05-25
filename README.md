# gomm

This project implements a library that compares two sorted files and tells which lines exist in only one of the files or exit in both. The functionality is the same as the [comm utility](https://en.wikipedia.org/wiki/Comm).

For example, to print only strings that appear in both lists:

```go
left := gomm.NewMemoryScanner([]string{"1", "2", "3"})
right := gomm.NewMemoryScanner([]string{"2", "3", "4"})

callback := func(p gomm.Position, line string) {
    if p == gomm.BOTH {
        fmt.Println(line)
    }
}

g := gomm.New(left, right, callback)
_ = g.Compare()
// Output:
// 2
// 3
```