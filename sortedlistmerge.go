package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1==nil{
		return n2
	}
	if n2==nil{
		return n1
	}
	data1:=n2.Data
	n1=	SortListInsert(n1, data1)
	n1=SortedListMerge(n1, n2.Next)

	return n1
}