package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	l1 := &ListNode{
		Val: 2,
	}
	l1.Add(&ListNode{
		Val: 4,
	})
	l1.Add(&ListNode{
		Val: 3,
	})

	l2 := &ListNode{
		Val: 5,
	}
	l2.Add(&ListNode{
		Val: 6,
	})
	l2.Add(&ListNode{
		Val: 4,
	})

	// l1 := &ListNode{
	// 	Val: 1,
	// }
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 0,
	// })
	// l1.Add(&ListNode{
	// 	Val: 1,
	// })

	// l2 := &ListNode{
	// 	Val: 5,
	// }
	// l2.Add(&ListNode{
	// 	Val: 6,
	// })
	// l2.Add(&ListNode{
	// 	Val: 4,
	// })
	fmt.Println(f4(l1, l2))

}

// this fails
func f(l1 *ListNode, l2 *ListNode) *ListNode {
	a := []int{}
	for l1 != nil {
		a = append(a, l1.Val)
		l1 = l1.Next
	}

	b := []int{}
	for l2 != nil {
		b = append(b, l2.Val)
		l2 = l2.Next
	}

	va := 0
	for i := len(a) - 1; i >= 0; i-- {
		va = va + int(math.Pow(10, float64(i)))*a[i]
	}

	vb := 0
	for i := len(b) - 1; i >= 0; i-- {
		vb = vb + int(math.Pow(10, float64(i)))*b[i]
	}

	v := va + vb
	str := strconv.FormatInt(int64(v), 10)
	d, _ := strconv.ParseInt(string(str[0]), 10, 64)
	node := &ListNode{
		Val: int(d),
	}
	for i := 1; i < len(str); i++ {
		d, _ := strconv.ParseInt(string(str[i]), 10, 64)
		n := &ListNode{
			Val:  int(d),
			Next: node,
		}
		node = n

	}
	return node
}

// this succeed
func f2(l1 *ListNode, l2 *ListNode) *ListNode {
	a1 := []int{}
	for l1 != nil {
		a1 = append(a1, l1.Val)
		l1 = l1.Next
	}

	a2 := []int{}
	for l2 != nil {
		a2 = append(a2, l2.Val)
		l2 = l2.Next
	}

	size := len(a1)
	if len(a1) < len(a2) {
		size = len(a2)
	}

	res := make([]int, 0, size+1)
	i := 0
	j := 0
	l := 0
	for i < len(a1) || j < len(a2) {
		v1 := 0
		if i < len(a1) {
			v1 = a1[i]
			i++
		}
		v2 := 0
		if j < len(a2) {
			v2 = a2[j]
			j++
		}

		v := v1 + v2 + l

		if v >= 10 {
			l = 1
			v = v - 10
		} else {
			l = 0
		}
		res = append(res, v)
	}
	if l > 0 {
		res = append(res, l)
	}
	current := &ListNode{
		Val: res[0],
	}
	head := current
	for i := 1; i < len(res); i++ {
		current.Next = &ListNode{
			Val: res[i],
		}
		current = current.Next
	}
	return head
}

func f3(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2
	l := 0

	var head, current *ListNode

	for c1 != nil || c2 != nil {
		v1 := 0
		if c1 != nil {
			v1 = c1.Val
			c1 = c1.Next
		}
		v2 := 0
		if c2 != nil {
			v2 = c2.Val
			c2 = c2.Next
		}

		v := v1 + v2 + l
		if v >= 10 {
			l = 1
			v = v - 10
		} else {
			l = 0
		}
		if head == nil {
			head = &ListNode{
				Val: v,
			}
			current = head
		} else {
			current.Next = &ListNode{
				Val: v,
			}
			current = current.Next
		}
	}
	if l > 0 {
		current.Next = &ListNode{
			Val: l,
		}
	}

	return head
}

func f4(l1 *ListNode, l2 *ListNode) *ListNode {
	nopHead := &ListNode{
		Val: 0,
	}
	current := nopHead
	carry := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		v := v1 + v2 + carry
		if v >= 10 {
			carry = 1
			v = v - 10
		} else {
			carry = 0
		}

		current.Next = &ListNode{
			Val: v,
		}
		current = current.Next
	}
	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}
	return nopHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
	Last *ListNode
}

func (n *ListNode) Add(node *ListNode) {
	c := n
	if c.Last == nil {
		c.Next = node
		c.Last = node
		return
	}
	c.Last.Add(node)
}
