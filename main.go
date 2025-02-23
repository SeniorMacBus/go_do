package main

import (
	"fmt"
	"go_do/rope"
)

func main() {
	r := rope.CreateNewRope("Alma", nil, nil)
	fmt.Println(r.Weight)
	fmt.Println(len("Alma"))
	fmt.Println(r.Data)
}
