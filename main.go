package main

import (
	"fmt"
	"reflect"

	"github.com/sergeysuvorov/go-btree/pkg/binarysearchtree"
)

func main() {
	values := []int{11, 12, 9, 8, 1, 20, 1}
	tree := binarysearchtree.NewBinarySearchTree(10)

	for _, v := range values {
		if err := tree.Insert(v); err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(reflect.TypeOf(tree))
	tree.PrintInOrder()

	searchVal := 3

	if _, ok := tree.Search(searchVal); ok {
		fmt.Printf("\n%d exists\n", searchVal)
	} else {
		fmt.Printf("\n%d not exists\n", searchVal)
	}

	fmt.Println("Min Value: ", tree.Min())
	fmt.Println("Max Value: ", tree.Max())
	fmt.Println("Total nodes: ", tree.CountNodes())

}
