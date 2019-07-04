package main

import "fmt"

// Stack is a structure to hold data
type Stack struct {
	top   bool
	value interface{}
	prev  *Stack
}

// New to create a new Stack
func (stk Stack) New() *Stack {
	return &stk
}

// Push is to push elements to stack
func (stk *Stack) Push(x interface{}) *Stack {
	if stk.top == false && stk.prev == nil {
		*stk = Stack{true, x, nil}
		return stk
	}
	if stk.top == true {
		// fmt.Println("stk=", stk)
		// fmt.Println("value=", x)
		stk.top = false
		val := stk.value
		top := stk.top
		prev := stk.prev
		*stk = Stack{true, x, &Stack{top, val, prev}}
		return stk
	}
	return stk
}

// Pop the element of Stack
func (stk *Stack) Pop() *Stack {
	if stk.top == false && stk.prev == nil {
		fmt.Println("stack is empty")
		return nil
	}
	if stk.prev == nil {
		fmt.Println("poped element= ", stk.value)
		fmt.Println("No more element to Pop")
		return stk
	}
	fmt.Println("poped element= ", stk.value)
	*stk = *stk.prev
	return stk
}

// TopElement is the value of Top
func (stk Stack) TopElement() interface{} {
	if stk.top == false && stk.prev == nil {
		fmt.Println("stack is empty")
		return nil
	}
	return stk.value
}

// PrintStack to list the elemet of stack
func (stk *Stack) PrintStack() {
	// fmt.Println("stack is ")

	if stk.top == false && stk.prev == nil {
		fmt.Println("stack is empty")
	}

	var res []interface{}
	res = append(res, stk.value)
	s := *stk
	for {
		s = *s.prev
		res = append(res, s.value)
		if s.prev == nil {
			break
		}
	}
	fmt.Println(res)
}

// Find to list the element of stack
func (stk *Stack) Find(findValue interface{}) {
	// fmt.Println("stack is ")
	var findflag bool

	if stk.top == false && stk.prev == nil {
		fmt.Println("stack is empty")
	}

	if findValue == stk.value {
		findflag = true
	}
	s := *stk
	if findflag == false {
		for {

			s = *s.prev
			if findValue == s.value {
				findflag = true
				break
			} else {
				findflag = false
			}
			if s.prev == nil {
				break
			}
		}
	}
	if findflag == true {
		fmt.Println("match found !!! = ", findValue)
	} else {
		fmt.Println("match not found !!!")
	}

}
