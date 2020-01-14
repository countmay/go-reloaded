package student 

type NodeI struct{
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI{
	n:=&NodeI{Data: data_ref}
	if l==nil{
		return n
	}
	prev:=l
	current:=l 
	count:=0
	for current.Data<data_ref{
		prev=current
		if current.Next!=nil{
			current=current.Next
			count++
		}else if current.Next==nil{
			current.Next=n
			return l
		}
	}
	if count>0{
		prev.Next=n 
		n.Next=current
		return l
	}
	n.Next=current
	return n
}
