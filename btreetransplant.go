package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}

	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root

	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root

}

func BTreeSearchItem(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return root
	}
	if root.Data == data {

		return root
	} else if data < root.Data {
		return BTreeSearchItem(root.Left, data)
	} else {
		return BTreeSearchItem(root.Right, data)
	}

}
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)

}
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	if node.Parent == nil {
		node = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else if node == node.Parent.Right {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root

}
