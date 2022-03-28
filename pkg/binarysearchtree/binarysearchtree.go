package binarysearchtree

import (
	"errors"
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

func NewBinarySearchTree(value int) *node {
	return &node{value: value}
}

func (n *node) Insert(value int) error {
	if n == nil {
		return errors.New("tree not exists")
	}

	if n.value == value {
		return fmt.Errorf("value %d already exists", value)
	}

	if n.value > value {
		if n.left == nil {
			n.left = &node{value: value}
			return nil
		}
		return n.left.Insert(value)
	}

	if n.value < value {
		if n.right == nil {
			n.right = &node{value: value}
			return nil
		}
		return n.right.Insert(value)
	}

	return nil
}

func (n *node) Search(value int) (node, bool) {
	if n == nil {
		return node{}, false
	}

	switch {
	case n.value == value:
		return *n, true
	case n.value > value:
		return n.left.Search(value)
	case n.value < value:
		return n.right.Search(value)
	}

	return node{}, false
}

func (n *node) PrintInOrder() {
	if n == nil {
		return
	}

	n.left.PrintInOrder()
	fmt.Printf("%d ", n.value)
	n.right.PrintInOrder()
}

func (n *node) Min() int {
	if n.left == nil {
		return n.value
	}

	return n.left.Min()
}

func (n *node) Max() int {
	if n.right == nil {
		return n.value
	}

	return n.right.Max()
}

func (n *node) CountNodes() int {
	if n == nil {
		return 0
	}

	return n.count(1)
}

func (n *node) count(carry int) int {

	if n.left != nil {
		carry = n.left.count(carry + 1)
	}

	if n.right != nil {
		carry = n.right.count(carry + 1)
	}

	return carry
}
