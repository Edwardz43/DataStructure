package main

import "fmt"

// Node is a LinkedList Node.
type Node struct {
	val  int
	next *Node
}

// List is a LinkedList.
type List struct {
	head *Node
	tail *Node
	len  int
}

// PreviousNode return the previous node of the specific pos.
func (l *List) PreviousNode(pos int) *Node {
	ptr := l.head

	if pos < 2 {
		return l.head
	}

	if pos > l.len {
		return l.tail
	}

	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}

	return ptr
}

// InsertAfter inserts a new node after the node of given pos.
func (l *List) InsertAfter(val int, ptr *Node) {
	tmpNode := Node{
		val:  val,
		next: ptr.next,
	}
	ptr.next = &tmpNode

	if l.tail == ptr {
		l.tail = &tmpNode
	}

	l.len++
}

// GetNode return the node of pos.
func (l *List) GetNode(pos int) *Node {

	if pos < 1 || pos > l.len {
		return nil
	}

	ptr := l.head

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr
}

// Insert inserts a new node to tail.
func (l *List) Insert(val int) {
	if l.head == nil {
		tmpNode := Node{
			val:  val,
			next: nil,
		}
		l.head = &tmpNode
		l.tail = &tmpNode
		l.len++
		return
	}

	l.InsertAfter(val, l.tail)
}

// Search return the node position of given value.
func (l *List) Search(val int) int {

	if l.len == 0 {
		return -1
	}

	ptr := l.head

	for i := 0; ptr != nil; i++ {
		if val == ptr.val {
			return i
		}
		ptr = ptr.next
	}

	return -1
}

// DeleteAt deletes the given position node.
func (l *List) DeleteAt(pos int) bool {
	if pos < 0 || pos > l.len-1 {
		return false
	}

	if pos == 0 {
		l.head = l.head.next
		l.len--
		return true
	}

	pre := l.PreviousNode(pos)

	pre.next = pre.next.next

	l.len--

	return true
}

// DeleteVal deletes the first node found with given value.
func (l *List) DeleteVal(val int) bool {
	pos := l.Search(val)
	return l.DeleteAt(pos)
}

// GetAt returns the node's value with given position.
func (l *List) GetAt(pos int) int {
	n := l.GetNode(pos)

	if n == nil {
		return -1
	}

	return n.val
}

func main() {
	list := new(List)

	arr := []int{1, 4, 5, 12, 75, 323}

	for _, v := range arr {
		list.Insert(v)
	}

	fmt.Println(list.len)
}
