package student

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {

	level := BTreeLevelCount(root)
	for i := 0; i < level; i++ {
		LevelMove(root, f, i)
	}

}

func LevelMove(root *TreeNode, f func(...interface{}) (int, error), lv int) {

	if root == nil {
		return
	}
	if lv == 0 {
		f(root.Data)
	} else if lv > 0 {
		LevelMove(root.Left, f, lv-1)
		LevelMove(root.Right, f, lv-1)
	}
}

func BTreeLevelCount(root *TreeNode) int {

	if root == nil {
		return 0
	}
	l := BTreeLevelCount(root.Left)
	r := BTreeLevelCount(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}
