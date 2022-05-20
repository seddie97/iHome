package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
type ListNode struct {
	Val  int
	Next *ListNode
}

//1,2,3,4,5 2
//1,2 2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	fast := head
	slow := &ListNode{Val: 0, Next: head}
	for n > 0 {
		n--
		fast = fast.Next
		if n != 0 && fast == nil {
			return head
		}
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return head
}

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}


func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
*/
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
