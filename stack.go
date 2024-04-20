package main

import "fmt"

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func main() {
	s := Stack{}
	s.Push(100)
	s.Push("hello world")
	s.Push("Kizaru")
	fmt.Println(s)
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println(s)
}
