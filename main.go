package main

import (
	"fmt"
	"log"
)

type Queue struct {
	arr []int
}

type Node struct {
	value int
	Next  *Node
}

func main() {
	head := &Node{2, &Node{ //1 task
		7, &Node{
			11, &Node{
				78, nil}}}}

	printReversedLinkedList(head)

	Q := Queue{} //2 task

	Q.push(1)
	Q.push(2)
	Q.push(3)

	fmt.Println(Q.arr)

	el, ok := Q.pop()
	if !ok {
		log.Fatal("Все нахуй, закончилось все")
	}
	fmt.Println("Съебалось:", el)
	fmt.Println(Q.arr)

	fmt.Println(isValid("(()){{}}")) // 3 task

	list1 := []int{1, 11, 5, 4}
	list2 := []int{17, 55, 6}
	list3 := []int{0}

	soartedList1 := sorting(list1)
	soartedList2 := sorting(list2)
	soartedList3 := sorting(list3)
	fmt.Println(merge(soartedList1, soartedList2, soartedList3)) //task4

	Q1 := Queue1{} //extra task 1
	Q1.frontPush(2)
	Q1.frontPush(14)
	Q1.frontPush(222)
	fmt.Println(Q1.arr)
	Q1.backPop()
	fmt.Println(Q1.arr)
	Q1.backPush(34554)
	fmt.Println(Q1.arr)
	Q1.frontPop()
	fmt.Println(Q1.arr)
}

func (q *Queue) push(v int) {
	q.arr = append(q.arr, v)
}

func (q *Queue) pop() (int, bool) {
	if len(q.arr) == 0 {
		return 0, false
	}

	v := q.arr[0]
	q.arr = q.arr[1:]

	return v, true
}

func printReversedLinkedList(head *Node) {
	var result []int

	for i := head; i != nil; i = i.Next {
		result = append(result, i.value)
	}

	fmt.Println(result)
}

func isValid(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, v := range s {

		if i, f := brackets[v]; f {
			if len(stack) == 0 || stack[len(stack)-1] != i {

				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}

func sorting(list []int) []int {
	n := len(list)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func merge(list1 []int, list2 []int, list3 []int) []int {
	merged := append(append(list1, list2...), list3...)
	sorting(merged)
	return merged
}

type Queue1 struct {
	arr []int
}

func (d *Queue1) frontPush(v int) {
	d.arr = append(d.arr, v)
}

func (d *Queue1) backPop() (int, bool) {
	if len(d.arr) == 0 {

		return 0, false
	}
	v := d.arr[len(d.arr)-1]
	d.arr = d.arr[:len(d.arr)-1]

	return v, true
}

func (d *Queue1) backPush(v int) {
	if len(d.arr) == 0 {
		d.arr = []int{v}
	} else {
		d.arr = append([]int{v}, d.arr...)
	}
}

func (d *Queue1) frontPop() (int, bool) {
	if len(d.arr) == 0 {

		return 0, false
	}

	v := d.arr[0]
	d.arr = d.arr[1:]

	return v, true
}
