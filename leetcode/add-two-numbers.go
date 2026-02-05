// https://leetcode.com/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    mainList := &ListNode{}
    current := mainList
    carryOver := 0

    for l1 != nil || l2 != nil || carryOver != 0 {
        sum := 0
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        total := sum + carryOver
        carryOver = total/10
        nodeVal := total%10

        current.Next = &ListNode{Val: nodeVal}
        current = current.Next
    }
    return mainList.Next
}