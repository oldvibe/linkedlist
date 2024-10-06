package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct{
	head *Node
	lenght int
}

func (list *LinkedList) insertAtHead(data int) {
	temp1 := &Node{data, nil}

	if list.head == nil {
		list.head = temp1
	} else  {
		temp2 := list.head
		list.head = temp1
		temp1.next = temp2
	}
	list.lenght += 1

}
func (list *LinkedList) display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ----> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}


func (l *LinkedList) insertAtTail(data int) {
    temp1 := &Node{data, nil}

    if l.head == nil {
        l.head = temp1
    } else {
        temp2 := l.head
        for temp2.next != nil {
            temp2 = temp2.next
        }
        temp2.next = temp1
    }
    l.lenght += 1
}


func (l *LinkedList) insert(n, data int) {
    if n == 0 {
        l.insertAtHead(data)
    } else if n == l.lenght-1 {
        l.insertAtTail(data)
    } else {
        temp1 := &Node{data, nil}
        temp2 := l.head

        for i := 0; i < n - 1; i++ {
            temp2 = temp2.next
        }
        temp1.next = temp2.next
        temp2.next = temp1
        l.lenght += 1
    }
}


func (l *LinkedList) deleteAtHead() {
    temp := l.head
    l.head = temp.next

    l.lenght -= 1
}

func (l *LinkedList) deleteAtTail() {
    temp1 := l.head
    var temp2 *Node
    for temp1.next != nil {
        temp2 = temp1
        temp1 = temp1.next
    }
    temp2.next = nil

    l.lenght -= 1
}

func (l *LinkedList) delete(n int) {
    if n == 0 {
        l.deleteAtHead()
    } else if n == l.lenght-1 {
        l.deleteAtTail()
    } else {
        temp1 := l.head
        for i := 0; i < n-1; i++ {
            temp1 = temp1.next
        }
        temp2 := temp1.next
        temp1.next = temp2.next
        l.lenght -= 1
    }
}

func (l *LinkedList) Reverse() {
    var curr, prev, next *Node
    curr = l.head
    prev = nil

    for curr != nil {
        next = curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    l.head = prev
}




func main() {
	l := &LinkedList{nil, 0}
	// first insert 
	l.insertAtHead(11)
	l.display()
	// second insert
	l.insertAtHead(10)
	l.display()
	//last insert
	l.insertAtHead(9)
	l.display()
	l.insertAtTail(7)
	l.display()

	l.insert(l.lenght - 1, 8)
	l.display()
	l.deleteAtHead()
	l.display()
	l.deleteAtTail()
	l.display()
	l.delete(2)
	l.display()
	l.Reverse()
	l.display()



}
