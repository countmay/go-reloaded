package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) *List {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return l
	}
	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	l.Tail = n
	return l
}
func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	if l.Tail == nil {
		if l.Head.Data == data_ref {
			l.Head = nil
		}
		return
	}
	prev := l.Head
	current := l.Head.Next
	if prev.Data == data_ref {
		l.Head = current
	}
	for current != l.Tail.Next {
		if current.Data == data_ref {
			current = current.Next
			prev.Next = current
			continue

		}
		prev = current
		current = current.Next

	}

}
