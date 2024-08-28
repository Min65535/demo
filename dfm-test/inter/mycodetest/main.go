package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//var reatNumMap = make(map[int]bool)

func deleteList(head *ListNode) *ListNode {
	fmt.Println("head val", head.Val)
	temp := head
	rep := new(ListNode)
	rep = nil
	for temp.Next != nil {
		val := temp.Val
		if temp.Val == temp.Next.Val {
			//reatNumMap[temp.Val] = true
			rep = &ListNode{
				Val:  val,
				Next: nil,
			}
			temp.Next = temp.Next.Next
		} else {
			if rep != nil {
				if rep.Val == val {
					temp.Val = temp.Next.Val
					temp.Next = temp.Next.Next
					rep = nil
				} else {
					temp = temp.Next
				}
			} else {
				temp = temp.Next
			}
			//if ok, _ := reatNumMap[val]; ok {
			//	temp.Val = temp.Next.Val
			//	temp.Next = temp.Next.Next
			//} else {
			//	temp = temp.Next
			//}
		}
	}
	fmt.Println("head val", head.Val)
	return head
}

func ex1() []int {
	h := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	temp := deleteList(h)
	var res []int
	for temp != nil {
		res = append(res, temp.Val)
		temp = temp.Next
	}
	return res
}

type Vec struct {
	By  int32
	Ind int
}

func ex2(s, t string) string {

	//var sm = make(map[*Vec]int)
	//var sl []*Vec
	//var tm = make(map[int32]int)
	//for i, v := range t {
	//	tm[v] = i
	//}
	//
	//for i, v := range s {
	//	if _, ok := tm[v]; ok {
	//		sm[&Vec{
	//			By:  v,
	//			Ind: i,
	//		}] = i
	//
	//		sl = append(sl, &Vec{
	//			By:  v,
	//			Ind: i,
	//		})
	//	}
	//}
	//
	//tmLength :=len(tm)
	//
	//for i,v := range sl{
	//	if
	//}
	return ""
}

func main() {
	fmt.Println(ex1())
	fmt.Println(ex2("ADOLBEAC", "ABC"))
}
