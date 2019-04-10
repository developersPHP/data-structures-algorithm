#第三章课后练习答案

1. 编写出打印一个单链表的所有元素程序
```
func PrintListElement(l *list) interface {
    cur := l.head
    for cur.next != nil {
        fmt.printft("node element is %v", cur.value)
        cur = cur.next
    }
}
```

2. 给你一个链表L和另一个链表P,它们包含升序排列的整数。操作PrintLots(L,P)将打印L中那些由P所指定的位置上的元素。例如，如果p=1，3，4，6，那么，L的第1，第3，第4和第6个元素被打印出来，写出过程PrintLots(L,P).你应该用基本的表操作，该过程运行的时间是多少
```
func PrintLots(L,P *LinkedList)  {
	cur := P.head.next
	for cur != nil {
		//遍历p位置链表上所有元素
		p := cur.value.(uint)
		if p >= L.length {
			fmt.Println("index is out of length")
			return 
		}
		curL := L.head.next
		var i uint = 0
		//遍历下标，找到对应元素
		for ; i< p; i++ {
			curL = curL.next
		}
		fmt.Printf("找到的元素为%v", curL.value)
		
		cur = cur.next
		
	}
}
时间复杂度为O(N)
```
3. 通过只调整指针(而不是数据)来交换两个相邻的元素，使用单链表或者双链表
```
func changeEachOther(p *LinkedList) {
	prev := p.head
	cur := p.head.next
	for cur != nil {
		//找出prev和next的node
		swap(prev, cur)
		prev = cur
		cur = cur.next
	}
}

func swap(prev, next *ListNode)  {
	var p1, p2 *ListNode
	p1 = prev
	p2 = next

	p1.next = p2.next
	p2.next = p1
}
```