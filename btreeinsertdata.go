package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	current := root

	for {
		if data < current.Data {
			if current.Left == nil {
				newNode := &TreeNode{Data: data}
				newNode.Parent = current
				current.Left = newNode
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				newNode := &TreeNode{Data: data}
				newNode.Parent = current
				current.Right = newNode
				break
			}
			current = current.Right
		}
	}
	return root
}
