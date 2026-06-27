/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    current := head 
	nodeMap := make(map[*ListNode]bool)

	for current != nil {
		_,exists := nodeMap[current] 
		if !exists {
			nodeMap[current] = true
			current = current.Next
		} else {
			return nodeMap[current]
		}
	}
	return false
}
