package list

import "fmt"

//单个结点的数据结构，用来表示一个结点的情况
type ListNode struct {
	next  *ListNode
	value interface{}
}

//链表数据结构，用来表示整个链表的情况
type LinkedList struct {
	head   *ListNode
	length uint
}

//创建一个结点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个结点p后面插入结点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		//如果在链表尾部插入结点，则不符合要求
		return false
	}
	newNode := NewListNode(v) //创建一个结点
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

//在某个结点p 前面插入结点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		//如果在链表头部加结点，或者指定结点为nil，都不符合要求
		return false
	}
	//获取头结点
	cur := this.head.next
	pre := this.head
	for nil != cur {
		//遍历链表,寻找到指定的结点
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		//指定的结点是尾结点
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true

}

//在链表头部插入结点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertBefore(this.head, v)
}

//在链表尾部插入结点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head      //获取链表头部
	for nil != cur.next { //遍历链表
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

//通过索引查找结点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的结点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil //启动gc垃圾回收结点
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
单链表反转
时间复杂度O(N)
*/
func (this *LinkedList) Reverse() {
	//链表为空，头结点为空，后继结点为空
	if nil == this.head || nil == this.head.next || nil == this.head.next.next {
		return
	}
	var pre *ListNode = nil
	cur := this.head.next
	for nil != cur {
		//交换两个结点，需要首先处理两个结点的指针指向
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	this.head.next = pre

}

/**
单链表是否有环
*/

func (this *LinkedList) HasCycle() bool {
	if nil != this.head {
		slow := this.head
		fast := this.head
		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/**
两个有序单链表合并
*/
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}

	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}

	l := &LinkedList{head: &ListNode{}}
	cur := l.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	for nil != cur1 && nil != cur2 {
		//如果两个链表不同长度，会首先把短的合并
		if cur1.value.(int) > cur2.value.(int) {
			cur.next = cur2
			cur2 = cur2.next
		} else {
			cur.next = cur1
			cur1 = cur1.next
		}
		cur = cur.next
	}
	//处理剩余的没有合并的链表
	if nil != cur1 {
		cur.next = cur1
	} else if nil != cur2 {
		cur.next = cur2
	}
	return l

}

func PrintLots(L, P *LinkedList) {
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
		for ; i < p; i++ {
			curL = curL.next
		}
		fmt.Printf("找到的元素为%v", curL.value)

		cur = cur.next

	}
}
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

func swap(prev, next *ListNode) {
	var p1, p2 *ListNode
	p1 = prev
	p2 = next

	p1.next = p2.next
	p2.next = p1
}
