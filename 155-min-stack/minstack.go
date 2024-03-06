package main

import "fmt"

type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {

	if len(m.min) == 0 || val <= m.GetMin() {
		m.min = append(m.min, val)
	}
	m.data = append(m.data, val)
}

func (m *MinStack) Pop() {

	if m.Top() == m.GetMin() {
		m.min = m.min[:len(m.min)-1]
	}

	m.data = m.data[:len(m.data)-1]
}

func (m *MinStack) Top() int {
	return m.data[len(m.data)-1]

}

func (m *MinStack) GetMin() int {

	return m.min[len(m.min)-1]

}

func main() {
	test := Constructor()

	test.Push(2)
	test.Push(0)
	test.Push(3)
	test.Push(0)
	fmt.Println(test.data)
	fmt.Println(test.GetMin())
	test.Pop()
	fmt.Println(test.data)
	fmt.Println(test.GetMin())
	test.Pop()
	fmt.Println(test.data)
	fmt.Println(test.GetMin())
	test.Pop()
	fmt.Println(test.data)
	fmt.Println(test.GetMin())
}
