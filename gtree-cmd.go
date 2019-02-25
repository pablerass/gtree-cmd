package main

import (
    "github.com/pablerass/gtree"
)

func main() {
    t := gtree.New(2)
    t.Insert(1)
    t.Insert(4)
    t.Print()
}
