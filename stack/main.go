package stack

import (
	"fmt"
	"errors"
)

type StackElement struct {
	Value interface{}
}

type Stack struct {
	Top int
	Size int
	Body []StackElement
}

type Stackable interface {
	IsEmpty() bool
	Peek() (*StackElement, error)
	Push(a interface{}) (*StackElement, error)
	Pop() (*StackElement, error)
}

func (self *Stack) Pprint() {
	fmt.Println("---TOP---")

	for i, e := range(self.Body) {
		fmt.Printf("%d - %s\r\n", i, e)
	}

	fmt.Println("---BOT---")
}

func (self *Stack) Find(a StackElement, compare func(a StackElement, b StackElement) bool) (int, bool) {
	for i, e := range(self.Body) {
		if compare(a, e) {			
			return i, true
		}
	}
	return 0, false
}

func (self *Stack) IsEmpty() bool {
	if self.Top == -1 {
		return true
	} else {
		return false
	}
}

func (self *Stack) Peek() (*StackElement, error) {
	if self.IsEmpty() {
		return nil, errors.New("empty")
	}
	return &self.Body[self.Top], nil
}

func (self *Stack) Push(a interface{}) (*StackElement, error) {
	newElement := StackElement{a}
	self.Top += 1
	if self.Top > (self.Size - 1) {
		return nil, errors.New("full")
	}
	self.Body[self.Top] = newElement
	return &newElement, nil
}

func (self *Stack) Pop() (*StackElement, error) {
	if self.Top >= 0 {
		v := self.Body[self.Top]		
		self.Top -= 1
		return &v, nil
	} else {
		return nil, errors.New("empty")
	}
}

func NewStack(size int) *Stack {
	body := make([]StackElement, size + 1)
	s := Stack{-1, size, body}
	return &s
}