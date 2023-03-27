package linkedlist

import "fmt"

type Node struct {
	value int
	prev *Node
	next * Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	length int
}

func (l *DoublyLinkedList) Insert(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head, l.tail = newNode, newNode
	}else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.length++
}

func (l *DoublyLinkedList) InsertAt(value, index int) {
	if index < 0 || index > l.length {
		fmt.Println("Err: invalid position")
		return
	}

	newNode := &Node{value: value}
	if index == 0 {
		if l.head == nil {
			l.head, l.tail = newNode, newNode
		}else {
			newNode.next = l.head
			l.head.prev = newNode
			l.head = newNode
		}
	}else if index == l.length {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}else {
		current := l.head
		for i := 1; i < index; i++ {
			current = current.next
		}
		newNode.next = current.next
		newNode.prev = current
		current.next.prev = newNode
		current.next = newNode
	}
	l.length++
}

func (l *DoublyLinkedList) Remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}else {
			l.tail = nil
		}
		l.length--
		return
	}

	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			if current.next != nil {
				current.next.prev = current
			}else {
				l.tail = current
			}
			l.length--
			return
		}
		current = current.next
	}
}

func (l *DoublyLinkedList) RemoveAt(index int) {
	if index < 0 || index >= l.length {
		fmt.Println("invalid position")
		return
	}

	if index == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}else {
			l.tail = nil
		}
	}else {
		current := l.head
		for i := 1; i < index; i++{
			current = current.next
		}
		current.next = current.next.next
		if current.next != nil {
			current.next.prev = current
		}else {
			l.tail = current
		}
	}
	l.length--
}

func (l *DoublyLinkedList) IndexOf(value int) int {
	current := l.head
	count := 0
	for current != nil {
		if current.value == value{
			return count
		}
		count ++
		current = current.next
	}
	return -1
}

func (l *DoublyLinkedList) Contains(value int) bool {
	current := l.head
	for current != nil {
		if current.value == value{
			return true
		}
		current = current.next
	}
	return false
}

func (l *DoublyLinkedList) Get(index int) (int, bool) {
	if index < 0 || index >= l.length {
		fmt.Println("invalid position")
		return 0, false
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, true
}

func (l *DoublyLinkedList) Reverse() {
	if l.head == nil || l.head.next == nil {
		return
	}

	current := l.head
	var temp *Node
	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	temp = l.head
	l.head = l.tail
	l.tail = temp
}

func (l *DoublyLinkedList) Clear() {
	l.head, l.tail, l.length = nil, nil, 0
}

func (l *DoublyLinkedList) Head() (int, bool) {
	if l.head == nil {
		fmt.Println("List is empty")
		return 0, false
	}
	return l.head.value, true
}

func (l *DoublyLinkedList) Tail() (int, bool) {
	if l.tail == nil {
		fmt.Println("List is empty")
		return 0, false
	}
	return l.tail.value, true
}

func (l *DoublyLinkedList) String() string {
	current := l.head
	var display string
	for current != nil {
		display += fmt.Sprintf("%d ->", current.value)
		current = current.next
	}
	display += "nil"
	return display
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l.length == 0
}

func (l *DoublyLinkedList) Size() int {
	return l.length
}

