func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	temp := head

	for temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return head
}