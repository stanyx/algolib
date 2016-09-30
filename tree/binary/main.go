package main

import (
	"fmt"
)

func GenerateTestTree() *BinTree {

	l31 := TNode{nil, nil, 5, nil}
	r32 := TNode{nil, nil, 10, nil}

	l33 := TNode{nil, nil, 16, nil}
	r34 := TNode{nil, nil, 18, nil}

	l21 := TNode{&l31, &r32, 9, nil}
	r22 := TNode{&l33, &r34, 17, nil}

	root := TNode{&l21, &r22, 15, nil}
	l21.P = &root
	r22.P = &root
	l31.P = &l21
	r32.P = &l21
	l33.P = &r22
	r34.P = &r22
	
	t := BinTree{&root, 3}

	return &t
}

type TNode struct {
	L *TNode
	R *TNode
	V interface{}
	P *TNode
}

func (self *TNode) Pprint() string {

	leftVal := -1
	rightVal := -1

	if self.L != nil {
		leftVal = self.L.V.(int)
	}

	if self.R != nil {
		rightVal = self.R.V.(int)
	}

	return fmt.Sprintf("TNode {%d} (left=%d, right=%d)", self.V, leftVal, rightVal)
}

type BinTree struct {
	Root *TNode
	H int
}

func (self *BinTree) Pprint() {

	self.InorderWalk(self.Root, 0, func(n *TNode, level int) {
		fmt.Printf("%d-%d | ", n.V.(int), level)
	})

}

func (self *BinTree) InorderWalk(n *TNode, level int, logFn func(n *TNode, level int)) {
	level += 1
	if n != nil {
		logFn(n, level)		
		if n.L != nil {
			self.InorderWalk(n.L, level, logFn)
		}
		if n.R != nil {
			self.InorderWalk(n.R, level, logFn)
		}
	}
}

func (self* BinTree) Search(node *TNode, v int) (*TNode) {

	if node == nil || v == node.V.(int) {
		return node
	}

	if v < node.V.(int) {
		return self.Search(node.L, v)
	} else {
		return self.Search(node.R, v)
	}

}

func (self* BinTree) IterativeSearch(node *TNode, v int) (*TNode) {

	for {
		if node != nil && node.V.(int) == v {
			return node
		}

		if v < node.V.(int) {
			node = node.L
		} else {
			node = node.R
		}

	}
}

func (self* BinTree) Max(x *TNode) (*TNode) {
	m := x
	for {
		if x == nil {
			break
		}
		m = x
		x = x.R
	}
	return m
}

func (self *BinTree) Min(x *TNode) (*TNode) {
	m := x
	for {
		if x == nil {
			break
		}
		m = x
		x = x.L
	}
	return m
}

func (self *BinTree) Successor(x *TNode) (*TNode) {

	if x.R != nil {
		return self.Min(x.R)
	}

	y := x.P

	for {
		if y != nil && x == y.R {
			x = y
			y = y.P
		} else {
			break
		}
	}

	return y
}

func (self *BinTree) Predessor(x *TNode) (*TNode) {

	if x.L != nil {
		return self.Max(x.R)
	}

	y := x.P

	for {
		if y != nil && x == y.L {
			x = y
			y = y.P
		} else {
			break
		}
	}

	return y
}

func (self* BinTree) Insert(z *TNode) {

	var y *TNode = nil
	x := self.Root

	for {
		if x == nil {
			break
		}
		y = x
		if z.V.(int) < x.V.(int) {
			x = x.L
		} else {
			x = x.R
		}
	}

	z.P = y

	if y == nil {
		self.Root = z
	} else if z.V.(int) < y.V.(int)  {
		y.L = z
	} else {
		y.R = z
	}
}

func (self *BinTree) Transplant(u *TNode, v *TNode) {

	if u.P == nil {
		self.Root = v
	} else if (u == u.P.L) {
		u.P.L = v
	} else {
		u.P.R = v
	}

	if v != nil {
		v.P = u.P
	}
}

func (self* BinTree) Delete(z *TNode) {
	if z.L == nil {
		self.Transplant(z, z.R)
	} else if z.R == nil {
		self.Transplant(z, z.L)
	} else {
		y := self.Min(z.R)
		if y.P != z {
			self.Transplant(y, y.R)
			y.R = z.R
			y.R.P = y
		}
		self.Transplant(z, y)
		y.L = z.L
		y.L.P = y
	}
}