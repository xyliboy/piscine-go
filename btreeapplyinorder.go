package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyInorder(root.Left, f) // 1) περνάμε πρώτα από τα ΑΡΙΣΤΕΡΑ παιδιά

	f(root.Data) // 2) εφαρμόζουμε τη συνάρτηση στο ΤΩΡΙΝΟ node

	BTreeApplyInorder(root.Right, f) // 3) στο τέλος πάμε στα ΔΕΞΙΑ παιδιά
}
