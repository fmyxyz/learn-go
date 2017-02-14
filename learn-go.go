package main

import (
	"fmt"
	"github.com/fmyxyz/learn-go/set"
)

func main() {
	mySet := set.NewHashSet()
	mySet.Add(1)
	mySet.Add(1)
	mySet.Add(2)
	fmt.Println(mySet)
}
