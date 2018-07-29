package main

import "fmt"

type Queue struct {
	elements []interface{}
	capacity int
	size     int
}

func NewQueue(capacity int) *Queue {
	q := new(Queue)
	q.elements = make([]interface{}, 0, capacity)
	q.capacity = capacity
	q.size = 0
	return q
}

func (q *Queue) Enqueue(element interface{}) {
	q.elements = append(q.elements, element)
	q.size++
}

func (q *Queue) Dequeue() interface{} { //出队
	if !q.IsEmpty() {
		front := q.elements[0]
		q.elements = q.elements[1:]
		q.size--
		return front
	}
	return nil
}

func (q *Queue) Peek() interface{} {
	if !q.IsEmpty() {
		return q.elements[0]
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func main() {
	q := NewQueue(10)

	fmt.Println(q.IsEmpty())

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	fmt.Println(q.elements)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Peek())

	for i := 0; i < 10; i++ {
		fmt.Println(q.Dequeue())
	}

	fmt.Println(q.elements)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Peek())

}
