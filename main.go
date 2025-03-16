package main

import (
	"fmt"
	"go_do/rope"
)

func main() {
	r1 := rope.CreateNewRope("Hello ", nil, nil)
	r2 := rope.CreateNewRope("Seaman!", nil, nil)
	r3 := rope.Concatenate(&r1, &r2)
	rope.Print(&r3)
	r4 := rope.CreateNewRope("Kakilnom kell", nil, nil)
	rope.Print(&r4)
	fmt.Println("")
	r5, r6, err := rope.Split(&r4, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	rope.Print(&r5)
	fmt.Println("")
	rope.Print(&r6)
}
