func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy

    p1, p2 := list1, list2

    // compare and merge
    for p1 != nil && p2 != nil {
        if p1.Val <= p2.Val {
            current.Next = p1
            p1 = p1.Next
        } else {
            current.Next = p2
            p2 = p2.Next
        }
        current = current.Next
    }

    if p1 != nil {
        current.Next = p1
    } else {
        current.Next = p2
    }

    return dummy.Next
}