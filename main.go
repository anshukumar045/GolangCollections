//Stack using single Linked List

package main

import "fmt"

func main() {

	// single linked list
  // Stack implementation
	var stk Stack
	a := stk.New()
	a.Push(10)
	a.Push(20)
	a.Push(30)
	a.Push("hello")
	fmt.Println(a.Find(10))
	a.PrintStack()
	fmt.Println(*a)
	a.Find("hello")
	a.Pop()
	fmt.Println("p1= ", a)
  a.Pop()
	fmt.Println("p1==", a)
	a.Pop()
	fmt.Println("p1==", a)
	a.Pop()
	fmt.Println("p1==", a)
}
