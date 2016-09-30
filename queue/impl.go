package queue

package queue

import (
	"fmt"
	"errors"
)

type QueueElement struct {
	Value interface{}
}

type Queue struct {
	Head int
	Tail int
	Size int
	Q []QueueElement
}

func NewQueue(size int) *Queue {
	body := make([]QueueElement, size)
	q := Queue{0, 0, size, body}
	return &q
}

func (self *Queue) InQueue() []QueueElement {
	q := []QueueElement{}
	offset := self.Head
	
	for {
		// fmt.Println(offset)
		// time.Sleep(time.Second)
		if self.Tail != offset {
			q = append(q, self.Q[offset])
		} else {
			break
		}

		if offset >= (self.Size - 1) {
			offset = 0
		} else {
			offset += 1
		}
	}

	return q
}

func (self *Queue) Enqueue(v interface{}) (*QueueElement, error) {

	if self.Head == self.Tail + 1 || (self.Head == 0 && self.Tail == self.Size - 1) {
		return nil, errors.New("full")
	}

	e := QueueElement{v}
	self.Q[self.Tail] = e

	if self.Tail >= (self.Size - 1) {
		self.Tail = 0
	} else {
		self.Tail += 1
	}
	
	return &e, nil
}

func (self *Queue) Dequeue() (*QueueElement, error) {

	if self.Head == self.Tail {
		return nil, errors.New("empty")
	}

	e := self.Q[self.Head]

	if self.Head >= (self.Size - 1) {
		self.Head = 0
	} else {
		self.Head += 1
	}

	return &e, nil
}

func (self *Queue) Pprint() {
	fmt.Printf("I ")
	for i, _ := range(self.Q) {
		fmt.Printf("%d    ", i)
	}

	fmt.Printf("\r\nV ")
	for _, e := range(self.Q) {
		if e.Value != nil {
			fmt.Printf("%d    ", e.Value)
		} else {
			fmt.Printf("%s", e)
		}
	}

	fmt.Println("\r\n")
}