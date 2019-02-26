package main

import (
    "fmt"
    "github.com/pablerass/gtree"
)

func main() {
    t := gtree.New("b", "2")
    t.Insert("c", "3")
    t.Insert("d", "4")
    t.Insert("a", "1")
    fmt.Println(t.Search("b"))
    t.Insert("b", "5")
    fmt.Println(t.Search("b"))
    fmt.Println(t.Search("e"))
    t.Print()
}
